import { createApp } from "vue";
import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import "./style.css";
import App from "./App.vue";
import axios from "axios";
import VueAxios from "vue-axios";
import router from "./router";
import store from "./store";
import { FontAwesomeIcon } from "./plugins/font-awesome";

const client = axios.create({
    baseURL: "/api/v1",
});

const app = createApp(App);
app.use(VueAxios, client);
app.use(router);
app.use(store);
app.component("font-awesome-icon", FontAwesomeIcon);
app.mount("#app");
