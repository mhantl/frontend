FROM node:alpine as buildera
WORKDIR '/usr/src/app'
COPY package.json .
#RUN npm install npm@7.11.1 --save react react-dom react-scripts
#RUN npm install react
#RUN npm install react-scripts
RUN npm install n lts
RUN npm install npm@7.11.2
RUN npm install react
RUN npm install react-scripts
COPY . .
RUN npm run build

FROM nginx
EXPOSE 80
COPY --from=buildera /usr/src/app/build /usr/share/nginx/html
