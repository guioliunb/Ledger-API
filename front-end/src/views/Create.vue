<template lang="pug">
    

    v-row

        v-col(md="6" offset-md="2")

            v-text-field(
                label="Name"
                v-model="name"
            )

            v-text-field(
                label="Weight"
                v-model="weight"
            )

            v-select(
                label="Type"
                v-model="type_id"
                :items="raw_resources"
                item-text="name"
                item-value="id"
            )

            v-btn(color="success" @click="create") Create

</template>
  
<script>

import { mapState } from 'vuex';

export default {
    name: "create",
    data(){
    return {
    };
    },
    computed:{
      ...mapState(
        {raw_resources: state => state.raw_resource_types}
      ),
    },
    methods: {
        async create(){
            const vm = this;

            await vm.$store.dispatch("create", {
                name: vm.name,
                weight: vm.weight,
                type_id: vm.type_id

            });

            await vm.$store.dispatch("fetchRawResourceTypes");

            vm.$router.push({
                path: "/"
            });

        }
    }
}
</script>
