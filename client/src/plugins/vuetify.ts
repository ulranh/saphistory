import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader
import Vue from 'vue';
import Vuetify from 'vuetify/lib';
// import light from './theme'

Vue.use(Vuetify);

export default new Vuetify({
  // theme: { light: true },
  icons: {
    iconfont: 'mdi', // default - only for display purposes
  },
  theme: {
    themes: {
      light: {
        background: '#EEEEEE',
      },
    },
  },
});
