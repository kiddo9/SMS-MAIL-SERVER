# Build frontend
FROM node:18 AS frontend-build

WORKDIR /app
COPY ../client/package*.json ./client/
RUN cd client && npm install
COPY ../client ./client
RUN cd client && npm run build

FROM nginx:alpine

# Copy NGINX config
COPY nginx.conf /etc/nginx/nginx.conf

# Copy React build
COPY --from=frontend-build /app/client/dist /usr/share/nginx/html

# Expose port
EXPOSE 80