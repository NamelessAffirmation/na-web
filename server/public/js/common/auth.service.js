import { ClientDataService } from "./client-data.service.js";

async function registerNewUser(email, username, password) {
  try {
    const response = await ClientDataService.post(
      "/api/auth/register",
      {
        email,
        username,
        password,
      },
      false
    );
    return response;
  } catch (error) {
    console.error(error.message);
    return null;
  }
}

async function sendLoginEmail(email) {
  try {
    const response = await ClientDataService.post(
      "/api/auth/send-login-email",
      {
        email,
      },
      false
    );
    return response;
  } catch (error) {
    console.error(error.message);
    return null;
  }
}

async function login(email, password) {
  try {
    const credentials = "Basic " + btoa(`${email}:${password}`);

    const response = await ClientDataService.authenticate(credentials);
    return response;
  } catch (error) {
    console.error(error.message);
    return null;
  }
}

export const AuthService = {
  registerNewUser,
  sendLoginEmail,
  login,
};
