<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>Account</h1>
        <v-btn @click="$router.push('/update')">Edit profile</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Account",
};
</script>

<style>
</style>