import '../styles/globals.css'
import 'bootstrap/dist/css/bootstrap.min.css'

function MyApp({ Component, pageProps }) {
  require('../i18n');
  return <Component {...pageProps} />
}

MyApp.getStaticProps = async (context) => {
  const { Component, ctx } = context
  return {...await Component.getStaticProps(ctx)}
}

export default MyApp
