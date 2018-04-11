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
        --url=${API_LOCAL_DOMAIN} \
        --title=${API_LOCAL_TITLE} \
        --admin_user=${API_LOCAL_ADMIN_USER} \
        --admin_password=${API_LOCAL_ADMIN_PASSWORD} \
        --admin_email=${API_LOCAL_ADMIN_EMAIL} \
        --skip-email
fi

# Update site title.
if [ "$API_LOCAL_TITLE" ]; then
    echo
    echo "Updating site title ..."
    vendor/bin/wp option update blogname "${API_LOCAL_TITLE}" --path=wordpress --allow-root
fi

# Update site description.
if [ "$API_LOCAL_DESCRIPTION" ]; then
    echo
    echo "Updating site description ..."
    vendor/bin/wp option update blogdescription "${API_LOCAL_DESCRIPTION}" --path=wordpress --allow-root
fi

# Checking out .org theme
if [ "tide" != "$API_LOCAL_THEME_NAME" ] && [ "twentyseventeen" != "$API_LOCAL_THEME_NAME" ]; then
    echo
    echo "Checking out theme ..."
    vendor/bin/wp theme install ${API_LOCAL_THEME_NAME} --path=wordpress --allow-root
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
if [ "$ACTIVE_THEME" != "${API_LOCAL_THEME_NAME}" ]; then
    echo
    echo "Activating theme ..."
    vendor/bin/wp theme activate ${API_LOCAL_THEME_NAME} --path=wordpress --allow-root
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
vendor/bin/wp user create audit-server audit-server@${API_LOCAL_DOMAIN} --role=audit_client --path=wordpress --allow-root

if [ "${TIDE_API_KEY}" ] && [ "${TIDE_API_SECRET}" ]; then
    echo
    echo "Setting audit-server credentials ..."
    vendor/bin/wp user meta update audit-server tide_api_user_key ${TIDE_API_KEY} --path=wordpress --allow-root
    vendor/bin/wp user meta update audit-server tide_api_user_secret ${TIDE_API_SECRET} --path=wordpress --allow-root
fi

echo
echo "Creating wporg user ..."
vendor/bin/wp user create wporg wporg@${API_LOCAL_DOMAIN} --role=api_client --path=wordpress --allow-root

echo
echo "Setup Complete!"

if [ -n "$(grep -i "$API_LOCAL_DOMAIN" /etc/hosts)" ]
    then
        echo "Found host"
    else
        echo "No host"
fi

echo
echo "Be sure to add '${BOLD}127.0.0.1 ${API_LOCAL_DOMAIN}${RESET}' to your hosts file."
echo "Use the following command (${RED}with caution${RESET}):"
echo
echo "${YELLOW}sudo echo \"127.0.0.1 ${API_LOCAL_DOMAIN}\" >> /private/etc/hosts${RESET}"
echo
exit 0
