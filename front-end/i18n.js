const NextI18Next = require('next-i18next').default
const { localeSubpaths } = require('next/config').default().publicRuntimeConfig
const path = require('path')

module.exports = new NextI18Next({
  otherLanguages: ['th'],
  localeSubpaths,
  localePath: path.resolve('./public/static/locales'),
  defaultLanguage: 'en',
  defaultNS: 'common'
})
