<template>
   <div class="columns">
    <router-link
      :to="{ name: 'Cobrar' }"
      class="button is-success mb-2"
    >
      Registrar nuevo puerp
    </router-link>
    <span>
      <b-field>
        <b-switch v-model="bars">Bars</b-switch>
      </b-field>
      <b-datepicker
        inline
        v-model="date"
        :events="events"
        :indicators="indicators"
        :unselectable-dates="unselectableDates"
      >
      </b-datepicker>
    </span>
    <section>
      <b-table :data="data" :columns="columns"> </b-table>
      <hr />

      
    </section>
  </div>
</template>
<script>

</script>