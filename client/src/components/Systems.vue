// html part
<template>
  <v-data-table
    dense
    fixed-header
    height="calc(100vh - 200px)"
    :headers="headers"
    :items="$store.state.systems"
    :sort-by="$store.state.systems"
    class="elevation-1"
    :footer-props="{
      itemsPerPageOptions: [20, -1],
    }"
  >
    <template v-slot:top>
      <v-toolbar flat color="white">
        <v-toolbar-title>SAP Systems</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-dialog v-model="dialog" max-width="600px">
          <template v-slot:activator="{ on }">
            <v-btn color="#577399" dark class="mb-2" v-on="on"
              >New system</v-btn
            >
          </template>
          <v-card>
            <v-card-title>
              <span class="headline">{{ formTitle }}</span>
            </v-card-title>

            <v-card-text>
              <v-container>
                <v-row>
                  <v-col cols="12" sm="6" md="6">
                    <v-text-field
                      v-model="editedItem.sid"
                      label="SID"
                      :rules="[rules.len3]"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" md="6">
                    <v-text-field
                      v-model="editedItem.client"
                      label="Client"
                      :rules="[rules.len3]"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" md="6">
                    <v-text-field
                      v-model="editedItem.description"
                      label="Description"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" md="6">
                    <v-text-field
                      v-model="editedItem.sysnr"
                      label="System Number"
                      :rules="[rules.len2]"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" md="6">
                    <v-text-field
                      v-model="editedItem.hostname"
                      label="Hostname"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" md="6">
                    <v-text-field
                      v-model="editedItem.username"
                      label="Username"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" md="6">
                    <v-text-field
                      v-model="editedItem.password"
                      label="Password"
                      :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                      :type="show ? 'text' : 'password'"
                      @click:append="show = !show"
                    ></v-text-field>
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
              <v-btn color="blue darken-1" text @click="save">Save</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-toolbar>
    </template>
    <template v-slot:item.actions="{ item }">
      <!-- <template> -->
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-icon
            small
            class="mr-2"
            v-on="on"
            v-bind="attrs"
            @click="gotoSystem(item.sid)"
            >mdi-exit-to-app</v-icon
          >
        </template>
        <span> Goto system data</span>
      </v-tooltip>
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-icon
            v-on="on"
            v-bind="attrs"
            small
            class="mr-2"
            @click="editItem(item)"
            >mdi-pencil</v-icon
          >
        </template>
        <span> Edit system</span>
      </v-tooltip>
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-icon
            v-on="on"
            v-bind="attrs"
            small
            class="mr-2"
            @click="deleteItem(item)"
            >mdi-delete</v-icon
          >
        </template>
        <span> Delete system</span>
      </v-tooltip>
    </template>
  </v-data-table>
</template>

// typescript part
<script lang="ts">
import { Component, Prop, Vue, Watch } from "vue-property-decorator";
import { Route } from "vue-router";
import {
  GetSystems,
  UpdateSystem,
  DeleteSystem,
} from "../internal/grpcFunctions";
import { SysInfo } from "../internal/system";

@Component
export default class SapSystems extends Vue {
  private headers = [
    {
      text: "SID",
      align: "start",
      sortable: "true",
      value: "sid",
    },
    {
      text: "Client",
      value: "client",
    },
    {
      text: "Description",
      value: "description",
    },
    {
      text: "System Nr.",
      value: "sysnr",
    },
    {
      text: "Hostname",
      value: "hostname",
    },
    {
      text: "Username",
      value: "username",
    },
    {
      text: "Actions",
      value: "actions",
      sortable: false,
    },
  ];

  // input rules
  private rules = {
    len3: (v: string) => v.length == 3 || "Exactly 3 characters",
    len2: (v: string) => v.length == 2 || "Exactly 2 characters",
  };

  private dialog = false;
  private show = false;
  private data: any = [];

  private editedIndex = -1;
  private editedItem: SysInfo = {
    sid: "",
    client: "",
    description: "",
    sysnr: "",
    hostname: "",
    username: "",
    password: "",
  };
  private defaultItem: SysInfo = {
    sid: "",
    client: "",
    description: "",
    sysnr: "",
    hostname: "",
    username: "",
    password: "",
  };

  // methods
  //
  private async mounted(): Promise<void> {
    // fetch systems data from database
    let data: any = await GetSystems(this.$store.state.grpcHost);
    this.$store.dispatch("setSystems", data);
  }

  // call dialog window for editing
  private editItem(item: SysInfo): void {
    item.password = "";
    this.editedIndex = this.$store.state.systems.indexOf(item);
    this.editedItem = Object.assign({}, item);
    this.dialog = true;
  }

  // delete system entry - frontend and backend
  private async deleteItem(item: SysInfo): Promise<void> {
    const index = this.$store.state.systems.indexOf(item);
    const system = this.$store.state.systems[index];
    confirm("Are you sure you want to delete this item?") &&
      this.$store.state.systems.splice(index, 1);
    const res = await DeleteSystem(this.$store.state.grpcHost, system);
  }

  // close input window
  private close(): void {
    this.dialog = false;
    setTimeout((): void => {
      this.editedItem = Object.assign({}, this.defaultItem);
      this.editedIndex = -1;
    }, 300);
  }

  // save system information - frontend and backend
  private async save(): Promise<void> {
    let val: any;
    if (this.editedIndex > -1) {
      Object.assign(
        this.$store.state.systems[this.editedIndex],
        this.editedItem
      );
      val = this.$store.state.systems[this.editedIndex];
    } else {
      this.$store.state.systems.push(this.editedItem);
      val = this.editedItem;
    }
    this.close();
    const res = await UpdateSystem(this.$store.state.grpcHost, val);
  }

  // jump to data page of this system
  private gotoSystem(sid: string): void {
    this.$store.dispatch("setSid", sid);
    this.$store.dispatch("setTitle", " - " + sid);
    this.$cookies.set("saphistorySid", sid);
    this.$router.push({
      name: "systemStatus",
      params: { sid: sid, ts: this.$store.state.ts, direction: "0" },
    });
  }

  //
  // computed properties
  //

  // correct title for input window
  get formTitle() {
    return this.editedIndex === -1 ? "New System" : "Edit System";
  }

  // watcher
  //
  @Watch("dialog")
  onDialogChanged(val: any) {
    val || this.close();
  }
}
</script>

// scss part
<style lang="scss">
/* .toolbar-title { */
/*   color: inherit; */
/*   text-decoration: inherit; */
/* } */
// .v-input__slot {
//   margin-bottom: 0 !important;
// }
</style>
