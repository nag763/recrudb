import { initializeApp } from "@firebase/app";
import { getAuth } from "@firebase/auth";

const gcpConfig = {
  apiKey: import.meta.env.VITE_GOOGLE_CLIENT_ID,
  authDomain: `${import.meta.env.PROJECT_NAME}.firebaseapp.com`,
  projectId: import.meta.env.PROJECT_NAME
};

const app = initializeApp(gcpConfig);

export const auth = getAuth(app);