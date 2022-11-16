<template>
  <div class="columns">
   
    <span>
      <h2 class="subtitle" >Citas medicas atrasadas MAIN </h2>
      
    </span>
    <section>
      <b-table :data="data" :columns="columns"> </b-table>
      <hr />

      
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
        first_name: "Jesse",
        last_name: "Simmons",
        date: "2016-10-15 13:43:27",
        sdg: "Semana 12",
      },
      {
        id: 2,
        first_name: "John",
        last_name: "Jacobs",
        date: "2016-12-15 06:00:53",
        sdg: "Semana 12",
      },
      {
        id: 3,
        first_name: "Tina",
        last_name: "Gilbert",
        date: "2016-04-26 06:26:28",
        sdg: "Semana 12",
      },
      {
        id: 4,
        first_name: "Clarence",
        last_name: "Flores",
        date: "2016-04-10 10:28:46",
        sdg: "Semana 12",
      },
      {
        id: 5,
        first_name: "Anne",
        last_name: "Lee",
        date: "2016-12-06 14:38:38",
        sdg: "Semana 12",
      },
    ],
    columns: [
      {
        field: "id",
        label: "No. Progreso",
        width: "100",
        numeric: true,
        searchable: true,
      },
      {
        field: "Nombre",
        label: "Nombre",
        searchable: true,
      },
      {
        field: "last_name",
        label: "Apellidos",
        searchable: true,
      },
      {
        field: "sdg",
        label: "SDG",
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