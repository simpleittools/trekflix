<template>
  <div class="form-container">
    <div class="head-container">
      <!--      replace with database logo file-->
      <img class="logo-image" src="../../assets/img/TrekFlix.png" alt="">
      <h2 class="page-title">Login</h2>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <form class="space-y-6" @submit.prevent="handleSubmit">
          <div class="mt-1 sm:mt-0 sm:col-span-2">
            <div class="input-frame">
              <span class="input-label"> Username </span>
              <input type="text" v-model="data.username" id="username" autocomplete="username" class="input-field" />
            </div>
          </div>

          <div class="mt-1 sm:mt-0 sm:col-span-2">
            <div class="input-frame">
              <span class="input-label"> Password </span>
              <input type="password" v-model="data.password" id="password" autocomplete="password" class="input-field" />
            </div>
          </div>
          <div class="button-container">
            <button type="submit" class="button-submit">Submit</button>
            <button type="reset" class="button-reset" @click.prevent="handleClear">Clear</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import {reactive} from "vue";

export default {
  name: "LoginView",
  setup() {
    const data = reactive({
      username: '',
      password: '',
    })

    const handleSubmit = () => {
      data.username
      data.password

      console.log(data)
    }
    const handleClear = () => {
      data.username = ''
      data.password = ''
      console.log(data)
    }

    return {
      data,
      handleSubmit,
      handleClear
    }
  }
}
</script>

<style scoped>

</style>