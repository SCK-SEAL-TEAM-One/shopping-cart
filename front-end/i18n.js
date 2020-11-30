import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import resources from './resources'

i18n
  .use(initReactI18next) // passes i18n down to react-i18next
  .init({
    resources: resources,
    lng: "th",
    fallbackLng: ["th", "en"],
    interpolation: {
      escapeValue: false
    },
  });



export default i18n;