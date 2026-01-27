import { initializeApp } from "@firebase/app";
import { getAuth } from "@firebase/auth";

const gcpConfig = {
  apiKey: import.meta.env.VITE_API_KEY,
  authDomain: `${import.meta.env.VITE_PROJECT_NAME}.firebaseapp.com`,
  projectId: import.meta.env.VITE_PROJECT_NAME
};

console.log(gcpConfig)

const app = initializeApp(gcpConfig);

export const auth = getAuth(app);