<template>
    <div>
      <title-bar :title-stack="titleStack" />
      <hero-bar>
        Tables
        <router-link
          slot="right"
          to="/"
          class="button"
        >
          Dashboard
        </router-link>
      </hero-bar>
      <section class="section is-main-section">

        <card-component
          class="has-table has-mobile-sort-spaced"
          title="Clients"
          icon="account-multiple"
        >
          <clients-table-sample
            checkable
          />
        </card-component>
  
        <hr>
  

  
        <card-component class="has-table has-mobile-sort-spaced">
          <clients-table-sample
            checkable
          />
        </card-component>
  
  
        <card-component class="has-table has-thead-hidden">
          <clients-table-sample is-empty />
        </card-component>
      </section>
    </div>
  </template>
  
  <script>

  import ClientsTableSample from '@/components/ClientsTableSample.vue'
  import CardComponent from '@/components/CardComponent.vue'
  import TitleBar from '@/components/TitleBar.vue'
  import HeroBar from '@/components/HeroBar.vue'
  
  export default {
    name: 'Tables',
    components: {
      HeroBar,
      TitleBar,
      CardComponent,
      ClientsTableSample,
    
    },
    data () {
      return {
        titleStack: ['Admin', 'Tables']
      }
    }
  }
  </script>