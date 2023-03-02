import { createRouter } from "vue-router";
import MainPagePage from "./components/MainPage.vue";

export default createRouter({
    routes: [
        {path: '/', component: MainPage}
    ]
})

