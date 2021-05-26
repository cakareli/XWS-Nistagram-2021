<template>
    <v-app>    
        <v-main>
            <h2>
                Hello Agent Guest!
            </h2>
            <router-view/>  
        </v-main> 
    </v-app>
</template>

<script>

export default {
  name: 'GuestHome',
}
</script>