<template>
  <div class="mt-1 sm:mt-0 sm:col-span-2">
    <div class="input-frame">
      <span class="input-label"> Last Name </span>
      <VeeField type="text" :value="modelValue"
             @input="$emit('update:modelValue', $event.target.value)"
             name="lastName" :rules="(value) => {
               if(value && value.trim()) return true
               return 'A Last Name is Required'
             }"

             autocomplete="lastName" class="input-field"
      />
    </div>
    <VeeErrorMessage name="lastName" class="form-error"/>
  </div>
</template>

<script>
import {ErrorMessage, Field} from "vee-validate";

export default{
  components: {
    VeeField: Field,
    VeeErrorMessage: ErrorMessage,

  },
  }
</script>

<script setup>
  defineProps({
    label: {
      type: String,
      default: "",
    },
    modelValue: {
      type: String,
      default: "",
    }
  })
</script>

<style scoped>

</style>