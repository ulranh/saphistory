// html part
<template>
  <div>
    <v-tabs
      v-model="tab"
      slider-color="#577399"
      color="#577399"
      background-color="#bdd5ea"
      grow
      align-with-title
    >
      <v-tab
        v-for="(tcTab, idx1) in result.tcvalues"
        :key="tcTab"
        :disabled="0 == result.data[idx1].length"
        ripple
        >{{ tcTab }}</v-tab
      >
    </v-tabs>
    <v-tabs-items v-model="tab" class="my-tabitem">
      <v-tab-item v-for="(tcTab, idx1) in result.tcvalues" :key="tcTab">
        <v-data-table
          dense
          fixed-header
          primary
          height="calc(100vh - 200px)"
          :headers="result.head[idx1]"
          :items="result.data[idx1]"
          class="elevation-1 mytable"
          :footer-props="{
            itemsPerPageOptions: [200, -1],
          }"
        >
        </v-data-table>
      </v-tab-item>
    </v-tabs-items>
  </div>
</template>

// typescript part
<script lang="ts">
import { Component, Prop, Vue, Watch } from "vue-property-decorator";
import { GetSapStatus } from "../internal/grpcFunctions";

@Component
export default class SystemStatus extends Vue {
  private tab = 0;
  private result: any = {};

  // private mounted() {
  // }

  // select new data from backend
  private async newData(
    host: string,
    sid: string,
    ts: string,
    direction: number
  ) {
    this.result = await GetSapStatus(host, sid, ts, direction);

    // adapt ts, prev ts and next ts
    await this.$store.dispatch("setTs", this.result.ts);

    await this.$store.dispatch(
      "setDate",
      this.result.ts.substr(0, 4) +
        "-" +
        this.result.ts.substr(4, 2) +
        "-" +
        this.result.ts.substr(6, 2)
    );
    await this.$store.dispatch(
      "setTime",
      this.result.ts.substr(8, 2) + ":" + this.result.ts.substr(10, 2)
    );
  }

  // watcher

  // select new data from backend, when route changes
  @Watch("$route", { immediate: true, deep: true })
  private onUrlChange() {
    this.newData(
      this.$store.state.grpcHost,
      this.$route.params.sid,
      this.$route.params.ts,
      Number(this.$route.params.direction)
    );
  }
}
</script>

// scss part
<style>
</style>