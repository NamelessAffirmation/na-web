function getClientDataConfiguration() {
  return fetch("/api/settings/client-data")
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    })
    .catch((error) => {
      console.error("There was a problem with the fetch operation:", error);
    });
}

function authenticate(credentials) {
  return fetch("/api/auth/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: credentials,
    },
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    })
    .catch((error) => {
      console.error("There was a problem with the fetch operation:", error);
    });
}

function post(url, body, isJSON) {
  return fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      if (isJSON) {
        return response.json();
      }
      return response.text();
    })
    .catch((error) => {
      console.error("There was a problem with the fetch operation:", error);
    });
}

function get(url, isJSON) {
  return fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      if (isJSON) {
        return response.json();
      }
      return response.text();
    })
    .catch((error) => {
      console.error("There was a problem with the fetch operation:", error);
    });
}

export const ClientDataService = {
  getClientDataConfiguration,
  authenticate,
  post,
  get,
};
