<template>
  <v-dialog>
    <template v-slot:activator="{ on }">
        <v-btn v-on="on">Register</v-btn>
    </template>
    <v-card>
        <v-card-title>
            <h2>Registration</h2>
        </v-card-title>
    </v-card>
    </v-dialog>
</template>

<script>
export default {
    name: 'Registration'
}
</script>

<style>

</style>