<template>
  <section>


    <b-table  class="table is-bordered is-striped is-narrow is-hoverable is-fullwidth" :data="data" >
      <template v-for="column in columns">
        <b-table-column .key="column.id" v-bind="column">
          <template 
          v-if="column.searchable && !column.numeric" 
          #searchable="props" >
            <b-input 
            v-model="props.filters[props.column.field]"
              placeholder="Buscar" 
              icon="magnify"
               size="is-small" />
          </template>
          <template v-slot="props">
            {{ props.row[column.field] }}
          </template>
        </b-table-column>
      </template>
    </b-table>
  </section>

</template>

<script>

export default {
  data() {
    return {
      data: [
        { 'id': 1, 'first_name': 'Jesse', 'last_name': '12', 'date': '2016-10-15 13:43:27', 'gender': 'Male' },
        { 'id': 2, 'first_name': 'John', 'last_name': '9', 'date': '2016-12-15 06:00:53', 'gender': 'Male' },
        { 'id': 3, 'first_name': 'Tina', 'last_name': '31', 'date': '2016-04-26 06:26:28', 'gender': 'Female' },
        { 'id': 4, 'first_name': 'Clarence', 'last_name': '20', 'date': '2016-04-10 10:28:46', 'gender': 'Male' },
        { 'id': 5, 'first_name': 'Anne', 'last_name': '4', 'date': '2016-12-06 14:38:38', 'gender': 'Female' }
      ],
      columns: [
        {
          field: 'id',
          label: 'No. Progreso',
          width: '100',
          numeric: true,
          searchable: true,
          centered: true
        },
        {
          field: 'first_name',
          label: 'Nombre completo',
          searchable: true,
          centered: true
        },
        {
          field: 'last_name',
          label: 'SDG',
          centered: true
        },
        {
          field: 'date',
          label: 'Fecha de consutal',
          searchable: true,
          centered: true
        }
      ]
    }
  }
}
</script>
<style>
#tabla {
  justify-content: center;
}
</style> 