FROM cypress/base:10
WORKDIR /app
# Copy our test page and test files
COPY cypress.json ./
COPY package.json ./
COPY cypress ./cypress
# Install npm dependencies, can also use "npm ci"
RUN npm install --save-dev cypress
RUN $(npm bin)/cypress verify
RUN $(npm bin)/cypress run