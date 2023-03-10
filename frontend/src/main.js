import Vue from 'vue'
import App from './App.vue'
import '@mdi/font/css/materialdesignicons.css'
import router from "./router"
import Utiles from "./services/Utiles"
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import store from './store'


Vue.use(Buefy);
Vue.config.productionTip = false



/* Fetch sample data */
store.dispatch('fetch', 'clients')

/* Default title tag */
const defaultDocumentTitle = 'SISMATERNA'

/* Collapse mobile aside menu on route change & set document title from route meta */
router.afterEach((to) => {
  store.commit('asideMobileStateToggle', false)

  if (to.meta && to.meta.title) {
    document.title = `${to.meta.title} — ${defaultDocumentTitle}`
  } else {
    document.title = defaultDocumentTitle
  }
})
// Filtros
//Vue.filter("formatearFecha", fechaComoCadena => Utiles.formatearFechaYHoraSegunLocale(new Date(fechaComoCadena)));
Vue.filter("minutosAHorasYMinutos", Utiles.minutosAHorasYMinutos);
Vue.filter("dinero", Utiles.formatearDinero);

new Vue({
  render: h => h(App),
  router,
  store,
}).$mount('#app')
