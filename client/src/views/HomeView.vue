<template>
  <div class="home">
    <h2>{{message}}</h2>
  </div>
</template>

<script>
export default {
  name: 'HomeView',
  components: {
  },
}
</script>

<script setup>
  import {useStore} from "vuex";
  import {ref, onBeforeMount} from "vue";
  import axios from "axios";


  const message = ref("You are not logged in")
  const store = useStore()

  onBeforeMount(async () =>{
    try {
      const {data} = await axios.get('auth/user')
      await store.dispatch('setAuth', true)

      message.value = `Hi ${data.first_name}`


    } catch(err) {
      await store.dispatch('setAuth', false)
    }

  })

</script>
