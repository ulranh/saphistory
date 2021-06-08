// html part
<template>
  <v-app
    id="app"
    :style="{ background: $vuetify.theme.themes.light.background }"
  >
    <v-app-bar app color="#577399" dark fixed>
      <v-toolbar-title class="display-1"
        >SAP History {{ $store.state.title }}</v-toolbar-title
      >
      <v-spacer></v-spacer>
      <template v-if="this.$route.name === 'systemStatus'">
        <v-row>
          <v-col cols="6" sm="5" md="3">
            <v-menu
              :close-on-content-click="false"
              v-model="menu1"
              :return-value.sync="date"
              transition="scale-transition"
              offset-y
              min-width="190px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="date"
                  prepend-icon="mdi-calendar-outline"
                  readonly
                  v-on="on"
                ></v-text-field>
              </template>
              <v-date-picker
                color="#577399"
                v-model="date"
                @click:date="
                  $refs.menu.save(date);
                  fetchData('0');
                  menu1 = false;
                "
                :allowedDates="allowedDates"
                @input="menu2 = false"
              ></v-date-picker>
            </v-menu>
          </v-col>
          <v-col cols="6" sm="5" md="3">
            <v-menu
              ref="menu"
              v-model="menu2"
              :close-on-content-click="false"
              :return-value.sync="time"
              transition="scale-transition"
              offset-y
              min-width="190px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="time"
                  prepend-icon="mdi-clock-outline"
                  readonly
                  v-on="on"
                ></v-text-field>
              </template>
              <v-time-picker
                color="#577399"
                v-if="menu2"
                v-model="time"
                format="24hr"
                @click:minute="
                  $refs.menu.save(time);
                  fetchData('0');
                "
              ></v-time-picker>
            </v-menu>
          </v-col>
          <v-col cols="6" sm="5" md="3">
            <span class="group pa-2">
              <v-tooltip bottom>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn icon @click="fetchData('-1')" v-bind="attrs" v-on="on">
                    <v-icon large>mdi-arrow-left-bold-box</v-icon>
                  </v-btn>
                </template>
                <span>Last scrape</span>
              </v-tooltip>
              <v-tooltip bottom>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn icon @click="fetchData('1')" v-bind="attrs" v-on="on">
                    <v-icon large>mdi-arrow-right-bold-box</v-icon>
                  </v-btn>
                </template>
                <span>Next scrape</span>
              </v-tooltip>
            </span>
          </v-col>
        </v-row>
        <v-spacer></v-spacer>
        <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      </template>
    </v-app-bar>
    <v-navigation-drawer v-model="drawer" absolute temporary right>
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title>SAP History</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list dense>
        <v-list-item
          v-for="item in items"
          :key="item.title"
          @click="gotoPage(item.route)"
          link
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>

// typescript part
<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import SystemStatus from "@/components/SystemStatus.vue";
import Systems from "@/components/Systems.vue";

@Component({
  components: {
    SystemStatus,
    Systems,
  },
})
export default class Root extends Vue {
  private menu1 = false;
  private menu2 = false;
  private drawer = null;
  private items = [
    { title: "Systems", icon: "mdi-desktop-classic", route: "systems" },
    //{ title: 'Preferences', icon: 'mdi-cog', route: 'preferences'},
    // { title: "About", icon: "mdi-forum" },
  ];

  // computed
  //

  // update values of datepicker and timepicker with the following getters and setters
  get date() {
    return this.$store.state.date;
  }
  set date(d) {
    this.$store.dispatch("setDate", d);
  }
  get time() {
    return this.$store.state.time;
  }
  set time(t) {
    this.$store.dispatch("setTime", t);
  }

  // methods
  //
  private mounted(): void {
    let host = {
      hostname: window.location.hostname,
      port: 8000, // initial port
    };

    // set port from environment, if VUE_APP_PORT exists
    // environment variables only work in vue, when starting with VUE_APP
    const port = Number(process.env.VUE_APP_PORT);
    if (port > 0) {
      host.port = port;
    }

    // decide between tls and non-tls
    // if VUE_APP_TLS_PATH "" or undefined -> http otherwise https
    if (
      process.env.VUE_APP_TLS_PATH !== "" &&
      process.env.VUE_APP_TLS_PATH !== undefined
    ) {
      this.$store.dispatch(
        "setGrpcHost",
        `https://${host.hostname}:${host.port}`
      );
    } else {
      this.$store.dispatch(
        "setGrpcHost",
        `http://${host.hostname}:${host.port}`
      );
    }

    // set initial date and time
    this.$store.dispatch("initDateTime");

    // if sap system sid is undefined
    // 1. look if system cookie is present
    // 2. otherwise goto systems page to select or create/select a system
    if ("" === this.$store.state.sid) {
      let cookieSid = this.$cookies.get("saphistorySid");
      if (cookieSid != undefined && cookieSid != "undefined") {
        this.$store.dispatch("setSid", cookieSid);
        this.$store.dispatch("setTitle", " - " + cookieSid);

        this.$router.push({
          name: "systemStatus",
          params: { sid: cookieSid, ts: this.$store.state.ts, direction: "0" },
        });
      } else {
        this.gotoPage("systems");
      }
    }
  }

  // define direction (one scrape forward or backward)
  private fetchData(direction: string) {
    const sid = this.$store.state.sid;
    const ts: string = this.getTs();

    // avoid error when same route is called more than once
    if (this.$route.path === `/systemStatus/${sid}/${ts}/${direction}`) return;

    this.$router.push({
      name: "systemStatus",
      params: { sid: this.$store.state.sid, ts: ts, direction: direction },
    });
  }

  // jump to page
  private gotoPage(route: string): void {
    this.$store.dispatch(
      "setTitle",
      " - " + route[0].toUpperCase() + route.slice(1)
    );
    this.$router.push({
      name: route,
      // params: {
      //   // port: this.$store.state.port,
      // },
    });
  }

  // return timestamp as string
  private getTs(): string {
    let ts =
      this.$store.state.date.substr(0, 4) +
      this.$store.state.date.substr(5, 2) +
      this.$store.state.date.substr(8, 2) +
      this.$store.state.time.substr(0, 2) +
      this.$store.state.time.substr(3, 2);

    return ts;
  }

  // return allowed dates for vuetify date/time picker
  private allowedDates(val: string) {
    return val <= this.$store.state.today;
  }

  // props
  //@Prop({required: true})
  //  dir:number = 0

  //@Watch($route)
  //onPropertyChanged(to: string, from: string) {
  //  document.title = to.meta.title || 'SAP History'
  //}
}
</script>

// scss part
<style lang="scss">
.v-input__slot {
  margin-bottom: 0 !important;
}
</style>
