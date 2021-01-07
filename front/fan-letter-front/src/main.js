import Vue from 'vue';
import App from './App.vue';
import vuetify from './plugins/vuetify';
import Axios from 'axios';

Vue.config.productionTip = false;

Vue.use(Axios);

Vue.config.productionTip = false;

new Vue({
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
