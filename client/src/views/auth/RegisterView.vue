<template>
  <div class="form-container">
    <div class="head-container">
<!--      replace with database logo file-->
      <img class="logo-image" src="../../assets/img/TrekFlix.png" alt="">
      <h2 class="page-title">Sign Up</h2>
      <p class="page-subtitle">
        to continue to TrekFlix
      </p>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <form class="space-y-6" @submit.prevent="handleSubmit">
          <div class="mt-1 sm:mt-0 sm:col-span-2">
            <div class="input-frame">
              <span class="input-label"> First Name </span>
              <input type="text" v-model="data.firstName" id="firstName" autocomplete="firstName" class="input-field" />
            </div>
          </div>
          <div class="mt-1 sm:mt-0 sm:col-span-2">
            <div class="input-frame">
              <span class="input-label"> Last Name </span>
              <input type="text" v-model="data.lastName" id="lastName" autocomplete="lastName" class="input-field" />
            </div>
          </div>
          <div class="mt-1 sm:mt-0 sm:col-span-2">
            <div class="input-frame">
              <span class="input-label"> Username </span>
              <input type="text" v-model="data.username" id="username" autocomplete="username" class="input-field" />
            </div>
          </div>
          <div class="mt-1 sm:mt-0 sm:col-span-2">
            <div class="input-frame">
              <span class="input-label"> Email </span>
              <input type="text" v-model="data.email" id="email" autocomplete="email" class="input-field" />
            </div>
          </div>
          <div class="mt-1 sm:mt-0 sm:col-span-2">
            <div class="input-frame">
              <span class="input-label"> Confirm Email </span>
              <input type="text" v-model="data.email2" id="email2" autocomplete="email2" class="input-field" />
            </div>
          </div>
          <div class="mt-1 sm:mt-0 sm:col-span-2">
            <div class="input-frame">
              <span class="input-label"> Password </span>
              <input type="password" v-model="data.password" id="password" autocomplete="password" class="input-field" />
            </div>
          </div>
          <div class="mt-1 sm:mt-0 sm:col-span-2">
            <div class="input-frame">
              <span class="input-label"> Verify Password </span>
              <input type="password" v-model="data.password2" id="password2" autocomplete="password2" class="input-field" />
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
  name: "RegisterView",
  setup() {
    const data = reactive({
      firstName: '',
      lastName: '',
      username: '',
      email: '',
      email2: '',
      password: '',
      password2: '',
    })

    const handleSubmit = () => {
      data.firstName
      data.lastName
      data.username
      data.email
      data.email2
      data.password
      data.password2

      console.log(data)
    }
    const handleClear = () => {
      data.firstName = ''
      data.lastName = ''
      data.username = ''
      data.email = ''
      data.email2 = ''
      data.password = ''
      data.password2 = ''

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