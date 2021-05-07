FROM node:12.6.0-alpine
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build
CMD npm start

# # Stage 2 - the production environment
# FROM nginx:1.16-alpine
# COPY --from=builder /app /usr/share/nginx/html
EXPOSE 3000
# CMD ["nginx", "-g", "daemon off;"]