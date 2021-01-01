import i18n from "i18next";
import { initReactI18next } from "react-i18next";
// import backend from "i18next-http-backend";

import common from "./common-messages.json";
import clients from './app/Clusters/messages.json'

const initOptions = {
  defaultNS: "common",
  resources: {
    en: {
      ...common,
      ...clients,
    },
  },
  lng: "en",
  fallbackLng: "en",

  interpolation: {
    escapeValue: false,
  },
};

i18n
  .use(initReactI18next)
  // .use(backend)
  .init(initOptions);

export default i18n;
