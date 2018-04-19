#!/bin/bash

## Output styles.
BOLD=$(tput bold)
RED=$(tput setaf 1)
YELLOW=$(tput setaf 3)
CYAN=$(tput setaf 6)
RESET=$(tput sgr0)

# Install Core
if ! $(vendor/bin/wp core is-installed --path=wordpress --allow-root); then
    echo
    echo "Installing WordPress ..."
    vendor/bin/wp core install \
        --allow-root \
        --path=wordpress \
        --url=${API_HTTP_HOST} \
        --title=${API_BLOG_NAME} \
        --admin_user=${API_ADMIN_USER} \
        --admin_password=${API_ADMIN_PASSWORD} \
        --admin_email=${API_ADMIN_EMAIL} \
        --skip-email
fi

# Update blogname.
if [ "${API_BLOG_NAME}" ]; then
    echo
    echo "Updating site title ..."
    vendor/bin/wp option update blogname "${API_BLOG_NAME}" --path=wordpress --allow-root
fi

# Update blogdescription.
if [ "${API_BLOG_DESCRIPTION}" ]; then
    echo
    echo "Updating site description ..."
    vendor/bin/wp option update blogdescription "${API_BLOG_DESCRIPTION}" --path=wordpress --allow-root
fi

# Checking out .org theme
if [ "tide" != "${API_THEME}" ] && [ "twentyseventeen" != "${API_THEME}" ]; then
    echo
    echo "Checking out theme ..."
    vendor/bin/wp theme install ${API_THEME} --path=wordpress --allow-root
fi

# Active theme
ACTIVE_THEME=`vendor/bin/wp theme list \
        --status=active \
        --fields=name \
        --format=csv \
        --allow-root \
        --path=wordpress \
        | tail -1 \
        2>/dev/null`

# Activating theme
if [ "${ACTIVE_THEME}" != "${API_THEME}" ]; then
    echo
    echo "Activating theme ..."
    vendor/bin/wp theme activate ${API_THEME} --path=wordpress --allow-root
fi

echo
echo "Activating plugins ..."

vendor/bin/wp plugin activate \
    flush-opcache \
    wp-tide-api \
    --path=wordpress \
    --allow-root

echo
echo "Setting up permalinks ..."
vendor/bin/wp rewrite structure '/%postname%' --path=wordpress --allow-root

echo
echo "Creating audit-server user ..."
vendor/bin/wp user create audit-server audit-server@${API_HTTP_HOST} --role=audit_client --path=wordpress --allow-root

if [ "${API_KEY}" ] && [ "${API_SECRET}" ]; then
    echo
    echo "Setting audit-server credentials ..."
    vendor/bin/wp user meta update audit-server tide_api_user_key ${API_KEY} --path=wordpress --allow-root
    vendor/bin/wp user meta update audit-server tide_api_user_secret ${API_SECRET} --path=wordpress --allow-root
fi

echo
echo "Creating wporg user ..."
vendor/bin/wp user create wporg wporg@${API_HTTP_HOST} --role=api_client --path=wordpress --allow-root

echo
echo "Setup Complete!"

echo
echo "Be sure to add '${BOLD}127.0.0.1 ${API_HTTP_HOST}${RESET}' to your hosts file."
echo "Use the following command (${RED}with caution${RESET}):"
echo
echo "${YELLOW}sudo echo \"127.0.0.1 ${API_HTTP_HOST}\" >> /private/etc/hosts${RESET}"
echo
exit 0
