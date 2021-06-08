import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
// import store from "../store/index";
import SystemStatus from "@/components/SystemStatus.vue";
import Systems from "@/components/Systems.vue";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "root",
    // component: Systems
  },
  {
    path: "/systemStatus/:sid/:ts/:direction",
    name: "systemStatus",
    component: SystemStatus,
    // props: castProps,
  },
  {
    // path: "/systems/:ts?",
    path: "/systems",
    name: "systems",
    component: Systems,
    // props: true,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});
export default router;
