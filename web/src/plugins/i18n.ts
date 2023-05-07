import {createI18n} from 'vue-i18n';
import enMessages from "../locales/en";
import zhMessages from "../locales/zh";

const supported = ['en', 'zh']
let locale = 'en'

try {
  const {0: browserLang} = navigator.language.split('-')
  console.log(browserLang)
  if (supported.includes(browserLang)) locale = browserLang
} catch (e) {
  console.log(e)
}

export const availableLocales = [
  {
    code: "en",
    flag: "united-states",
    label: "English",
    message: enMessages,
  },
  {
    code: "zh",
    flag: "china",
    label: "中文",
    message: zhMessages,
  },
]

const messages: {[key: string]: any} = {}
availableLocales.forEach(locale => {
  messages[locale.code] = locale.message
})

export default createI18n({
  legacy: false,
  locale: locale,
  fallbackLocale: 'zh',
  messages: messages,
});
