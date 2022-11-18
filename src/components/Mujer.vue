<template>
  <div class="columns" id="tabla">
  <section >
    <h2 class="title" >Citas medicas pendientes</h2>
    <b-table :data="data" :columns="columns"> </b-table>
      
    </section>
  </div>
  
</template>
<script>
const thisMonth = new Date().getMonth();
export default {
  computed: {
    indicators() {
      return this.bars ? "bars" : "dots";
      
    },
  },

  data: () => ({
    vehiculos: [],

    date: new Date(2022, thisMonth, 1),
    events: [
      new Date(2022, thisMonth, 2),
      new Date(2022, thisMonth, 6),
      {
        date: new Date(2022, thisMonth, 6),
        type: "is-info",
      },
      {
        date: new Date(2022, thisMonth, 8),
        type: "is-danger",
      },
      {
        date: new Date(2022, thisMonth, 10),
        type: "is-success",
      },
      {
        date: new Date(2022, thisMonth, 10),
        type: "is-link",
      },
      new Date(2022, thisMonth, 12),
      {
        date: new Date(2022, thisMonth, 12),
        type: "is-warning",
      },
      {
        date: new Date(2022, thisMonth, 16),
        type: "is-danger",
      },
      new Date(2022, thisMonth, 20),
      {
        date: new Date(2022, thisMonth, 29),
        type: "is-success",
      },
      {
        date: new Date(2022, thisMonth, 29),
        type: "is-warning",
      },
      {
        date: new Date(2022, thisMonth, 29),
        type: "is-info",
      },
    ],
    unselectableDates: [new Date(2022, thisMonth, 29)],
    bars: false,

    data: [
      {
        id: 1,
        Nombre: "Jesse",
        Apellidos: "Simmons",
        date: "2016-10-15 13:43:27",
        sdg: "10",
      },
      {
        id: 2,
        Nombre: "John",
        Apellidos: "Jacobs",
        date: "2016-12-15 06:00:53",
        sdg: "9",
      },
      {
        id: 3,
        Nombre: "Tina",
        Apellidos: "Gilbert",
        date: "2016-04-26 06:26:28",
        sdg: "6",
      },
      {
        id: 4,
        Nombre: "Clarence",
        Apellidos: "Flores",
        date: "2016-04-10 10:28:46",
        sdg: "2",
      },
      
    ],
    columns: [
      {
        field: "id",
        label: "No. Progreso",
        width: "100",
        placeholder:"Search",
        numeric: true,
        searchable: true,
        centered: true,
        
        
      },
      {
        field: "Nombre",
        label: "Nombre",
        placeholder:"Search",
        searchable: true ,
        centered: true,
        
      },
      {
        field: "Apellidos",
        label: "Apellidos",
        placeholder:"Search",
        searchable: true,
        centered:true,
      },
      {
        field: "sdg",
        label: "SDG",
        centered: true,
      },
      {
        field: "date",
        label: "Fecha de consulta",
        centered: true,
      },
     
    ],
  }),
  methods: {},
};
</script>
<style>
#tabla{
  justify-content: center;
}
</style>