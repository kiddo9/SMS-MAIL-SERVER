#set nuild context to root folder
# Build frontend
FROM node:22 AS frontend-build

WORKDIR /app
COPY ./client/package*.json ./client/
RUN cd client && npm install
COPY ./client ./client
RUN cd client && npm run build

FROM nginx:alpine

# Remove the default config
RUN rm /etc/nginx/conf.d/default.conf

# Copy NGINX config
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf

# Copy React build
COPY --from=frontend-build /app/client/dist /usr/share/nginx/html

# Expose port
EXPOSE 80