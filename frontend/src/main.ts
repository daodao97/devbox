import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import { regCustomFormComps } from '@okiss/vbtf'
import SliceForm from '@/components/SliceForm.vue'
import directives from './util/directives'

import '@/styles/index.scss'
import '@okiss/vbtf/style.css'

regCustomFormComps({ SliceForm })
createApp(App)
  .use(directives)
  .use(createPinia())
  .use(router)
  .mount('#app')
