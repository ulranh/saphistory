import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    systems: [],
    title: "",
    sid: "",
    today: "",
    date: "",
    time: "",
    ts: "",
    grpcHost: "",
    // direction: 0,
  },
  mutations: {
    grpcHost(state, host) {
      state.grpcHost = host;
    },
    systems(state, systems) {
      state.systems = systems;
    },
    // direction(state, direction) {
    //   state.direction = direction;
    // },
    sid(state, sid) {
      state.sid = sid;
    },
    title(state, title) {
      state.title = title;
    },
    today(state, today) {
      state.today = today;
    },
    ts(state, ts) {
      state.ts = ts;
    },
    date(state, date) {
      state.date = date;
    },
    time(state, time) {
      state.time = time;
    },
  },
  actions: {

    // set new grpc host url
    setGrpcHost(context, host) {
      context.commit("grpcHost", host);
    },

    // set new sap system
    setSystems(context, systems) {
      context.commit("systems", systems);
    },

    // 
    // setDirection(context, direction) {
    //   context.commit("direction", direction);
    // },

    // set new app-bar title
    setTitle(context, title) {
      context.commit("title", title);
    },
    setSid(context, sid) {
      context.commit("sid", sid);
    },

    // set initial date/time
    initDateTime(context) {
      let raw = new Date();
      let now = new Date(
        raw.getTime() - raw.getTimezoneOffset() * 60000
      ).toISOString();
      context.commit("today", now.substr(0, 10));
      context.commit("date", now.substr(0, 10));
      context.commit("time", now.substr(11, 5));
      context.commit(
        "ts",
        now.substr(0, 4) +
        now.substr(5, 2) +
        now.substr(8, 2) +
        now.substr(11, 2) +
        now.substr(14, 2)
      );
    },

    // set new date
    setDate(context, date) {
      context.commit("date", date);
    },

    // set new time
    setTime(context, time) {
      context.commit("time", time);
    },

    // set new timestamp
    setTs(context, ts) {
      context.commit("ts", ts);
    },
  },
  modules: {},
});
