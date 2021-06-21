<template>
  <v-app class="grey lighten-4">
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-toolbar height="45" width="800px">
            <v-app-bar app height="45" color="grey">
              <v-app-bar-nav-icon @click="$router.push('/account')">
                <v-icon>mdi-arrow-left</v-icon>
              </v-app-bar-nav-icon>
              <v-row>
                <v-col width="550px">
                  <v-toolbar-title>
                    <span outlined class="font-weight-light">NISTA</span>
                    <span>GRAM</span>
                  </v-toolbar-title>
                </v-col>
              </v-row>
            </v-app-bar>
          </v-toolbar>
          <v-container>
            <v-snackbar v-model="snackbar" top timeout="3500">
              <span>{{ snackbarText }}</span>
            </v-snackbar>
            <v-form ref="profilePrivacyForm">
              <v-row>
                <v-col>
                  <v-switch
                    class="mx-15"
                    @click="privacyTypeSwitch"
                    v-model="privacyTypeSW"
                    inset
                    :label="`Profile privacy: ${privacyTypeString.toString()}`"
                  ></v-switch>
                  <v-switch
                    class="mx-15"
                    @click="allMessageRequestsSwitch"
                    v-model="allMessageRequestsSW"
                    inset
                    :label="`Message requests: ${allMessageRequestsString.toString()}`"
                  ></v-switch>
                  <v-switch
                    class="mx-15"
                    @click="tagsSwitch"
                    v-model="tagsSW"
                    inset
                    :label="`Tags: ${tagsString.toString()}`"
                  ></v-switch>
                  <v-btn class="success" @click="updateProfilePrivacy"
                    >Apply changes</v-btn
                  >
                </v-col>
              </v-row>
            </v-form>
          </v-container>
        </v-col>
        <v-col width="300px"></v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<script>
import axios from "axios";
import { getId, getToken } from "../security/token.js";

export default {
  name: "ProfilePrivacy",

  mounted() {
    this.loadRegisteredUser();
  },

  data() {
    return {
      privacyTypeSW: "",
      privacyTypeString: "Public",
      allMessageRequestsSW: "",
      allMessageRequestsString: "Allowed",
      tagsSW: "",
      tagsString: "Allowed",
      snackbar: false,
      snackbarText: "",
      form: {},
    };
  },
  methods: {
    privacyTypeSwitch() {
      if (this.privacyTypeSW) {
        this.privacyTypeString = "Public";
      } else {
        this.privacyTypeString = "Private";
      }
    },

    allMessageRequestsSwitch() {
      if (this.allMessageRequestsSW) {
        this.allMessageRequestsString = "Allowed";
      } else {
        this.allMessageRequestsString = "Denied";
      }
    },

    tagsSwitch() {
      if (this.tagsSW) {
        this.tagsString = "Allowed";
      } else {
        this.tagsString = "Denied";
      }
    },

    updateProfilePrivacy() {
      if (this.$refs.profilePrivacyForm.validate()) {
        axios
          .put("http://localhost:8081/api/user/update-profile-privacy", {
            _id: getId(),
            privacyType: this.privacyTypeSW ? 0 : 1,
            allMessageRequests: this.allMessageRequestsSW ? true : false,
            tagsAllowed: this.tagsSW ? true : false,
          })
          .then((response) => {
            console.log(response.status);
            this.snackbarText =
              "Your profile privacy has been successfuly updated!";
            this.snackbar = true;
            setTimeout(() => {
              this.$router.push({ path: "/account" });
            }, 2000);
          })
          .catch((error) => {
            if (error.response.status === 400) {
              this.snackbarText = "Bad request!";
              this.snackbar = true;
            } else if (error.response.status === 409) {
              this.snackbarText =
                "User with same email address already exists!";
              this.snackbar = true;
            }
          });
      } else {
        this.snackbarText = "Please fill all fields correctly!";
        this.snackbar = true;
      }
    },

    loadRegisteredUser() {
      axios
        .get("http://localhost:8081/api/user/logged-user/" + getId(), {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.regularUser = response.data;
          if (this.regularUser.ProfilePrivacy.PrivacyType == 0) {
            this.privacyTypeSW = true;
          } else {
            this.privacyTypeSW = false;
          }

          this.allMessageRequestsSW = this.regularUser.ProfilePrivacy.AllMessageRequests;
          this.tagsSW = this.regularUser.ProfilePrivacy.TagsAllowed;

          this.privacyTypeSwitch();
          this.allMessageRequestsSwitch();
          this.tagsSwitch();
        });
    },
  },
};
</script>

