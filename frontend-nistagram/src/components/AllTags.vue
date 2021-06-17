<template>
    <v-dialog v-model="allTagsDialog" width = "500" persistent>
        <v-layout justify-center>
             <v-flex>
            <v-card class="pa-3">
                <v-card-title>
                        <h4>All tags:</h4>
                </v-card-title>
                    <span v-for="(tag, index) in allPostTags" :key ="index" class="ma-2">#{{tag}} </span>
                <v-card-actions>
                    <v-btn color="grey lighten-3" @click.native="close">Close</v-btn>
                </v-card-actions>
            </v-card>
             </v-flex>
        </v-layout>
         
    </v-dialog>
</template>
    
    
<script>

export default {
    name: 'AllTags',
    components: {

    },
    props: {
        allTagsDialog: {
        default: false,
      },
      allPostTags: {}
    },
    data() {
        return {

        }
    },
    mounted() {

    },
    methods: {
        close() {
            this.$emit('update:allTagsDialog', false)
        },
    }
}
</script>