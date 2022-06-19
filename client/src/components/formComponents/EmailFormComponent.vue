<template>
  <div class="mt-1 sm:mt-0 sm:col-span-2">
    <div class="input-frame">
      <span class="input-label"> Email </span>
      <VeeField type="email" :value="modelValue"
             @input="$emit('update:modelValue', $event.target.value)"
             name="email" :rules="(value) => {
               if(value && value.trim()) return true
               return 'An Email is Required'
             }"
                autocomplete="email" class="input-field"
      />
    </div>
    <VeeErrorMessage name="email" class="form-error"/>
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