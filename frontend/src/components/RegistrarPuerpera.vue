<template>
  <!-- <div class="columns"> -->
  <b-steps
    v-model="activeStep"
    :animated="isAnimated"
    :rounded="isRounded"
    :has-navigation="hasNavigation"
    :icon-prev="prevIcon"
    :icon-next="nextIcon"
    :mobile-mode="mobileMode"
  >
    <div class="has-text-centered">
      <article class="panel is-info centered">
        <div class="panel-heading">Registrar puerpera</div>
      </article>
    </div>
    <div>
      <div class="columns">
        <div class="column is-half">
          <b-field label="No. Progreso">
            <b-input
              disabled="»disabled»"
              v-model="detalles.id"
              placeholder=""
            ></b-input>
          </b-field>
          <b-field label="Nombre de puerpera">
            <b-input
              v-model="detalles.NombrePuerpera"
              placeholder="Nombre completo"
              disabled="»disabled»"
            ></b-input>
          </b-field>

          <b-field label=" Consulta de 7 días">
            <b-datepicker
              v-model="detalles.consultasiete"
              placeholder="Fecha consulta de 7 días"
              icon="calendar-today"
              :locale="locale"
              editable
            >
            </b-datepicker>
          </b-field>

          <b-field label="Consulta de 40 días">
            <b-datepicker
              v-model="detalles.consultacuarenta"
              placeholder="Fecha consulta de 40 días"
              icon="calendar-today"
              :locale="locale"
              editable
            >
            </b-datepicker>
          </b-field>

          <b-field label="Atención de parto">
            <b-select placeholder="¿Quién atendió el parto?" expanded>
              <option value="medico">Médico</option>
              <option value="emfermera">Enfermera</option>
              <option value="parteraTra">Partera Trad.</option>
              <option value="parteraPro">Partera Prof.</option>
              <option value="taps">TAPS</option>
              <option value="familiar">Familiar</option>
              <option value="ningunoAten">Ninguno</option>
            </b-select>
          </b-field>
          <b-field label="Resolución del embarazo">
            <b-select placeholder="Resolución del embarazo" expanded>
              <option value="vaginal">Vaginal</option>
              <option value="cesarea ">Cesarea</option>
              <option value="aborto">Legrado (aborto)</option>
              <option value="óbito">Legrado (óbito)</option>
              <option value="ningunoReso">Ninguno</option>
            </b-select>
          </b-field>

          <b-field label="APEO">
            <b-select placeholder="APEO" 
            v-model="detalles.APEO" expanded>
              <option value="diu">DIU</option>
              <option value="otb">OTB</option>
              <option value="implante">Implante</option>
              <option value="ningunoapeo">Ninguno</option>
            </b-select>
          </b-field>
          <b-field label="Consulta pregestacional"></b-field>
          <label class="radio">
            <input type="radio" name="pregestacional" />
            Si
          </label>
          <label class="radio">
            <input type="radio" name="pregestacional" />
            No
          </label>
          <b-field label="Derechohabiencia">
            <b-select placeholder="Derechohabiencia" disabled expanded>
              <option value="Insabi">INSABI</option>
              <option value="Imss">IMSS</option>
              <option value="Issste">ISSSTE</option>
              <option value="Sedena">SEDENA</option>
              <option value="Pemex">PEMEX</option>
              <option value="Marina">MARINA</option>
              <option value="Otro">Otro</option>
            </b-select>
          </b-field>
        </div>
        <div class="column is-half">
          <b-field label="CURP">
            <b-input
              v-model="detalles.descripcion"
              placeholder="Clave Unica de Registro de Población"
              icon="magnify"
            >
            </b-input>
          </b-field>
          <b-field label="Edad">
            <b-input disabled="»disabled»" v-model="detalles.Edad"></b-input>
          </b-field>

          <b-field label="Consulta de 28 días">
            <b-datepicker
              v-model="detalles.consultaveinte"
              placeholder="Fecha consulta de 28 días"
              icon="calendar-today"
              :locale="locale"
              editable
            >
            </b-datepicker>
          </b-field>
          <b-field label="Signos de alarma">
            <b-select placeholder="Signos" expanded>
              <option value="Fiebre">Fiebre</option>
              <option value="DolorPelvico">Dolor Pélvico</option>
              <option value="FlujoVaginal">Flujo Vaginal</option>
              <option value="FiebreDolor">Fiebre/Dolor Pelvico</option>
              <option value="FiebreFlujo">Fiebre/Flujo Vaginal</option>
              <option value=">DolorFlujo">Dolor Pelvico/Flujo Vaginal</option>
              <option value="FiebreDolorFlujo">
                Fiebre/Dolor Pelvico/Flujo Vaginal
              </option>
              <option value="Cefalea">Cefalea</option>
              <option value="NingunoSigno">Ninguno</option>
            </b-select>
          </b-field>

          <b-field label="Lugar del parto">
            <b-select placeholder="¿Dónde atendió el parto?" expanded>
              <option value="tlapa">H.G. Tlapa</option>
              <option value="hmnig">HMNIG</option>
              <option value="acatepec">H.C. Acatepec</option>
              <option value="alcozauca">H.C. Alcozauca</option>
              <option value="huamux">H.C. Huamuxtitlán</option>
              <option value="malina">H.C. Malinaltepec</option>
              <option value="oli">H.C. Olinalá</option>
              <option value="xochi">H.C. H.C. Xochihuehuetlán</option>
              <option value="zapo">H.C. Zapotitlán Tablas</option>
              <option value="centro">Centro de Salud</option>
              <option value="hogar">Hogar</option>
              <option value="otro">Otro hospital…</option>
            </b-select>
          </b-field>

          <b-field label="No. de producto">
            <b-select placeholder="No. de producto" expanded>
              <option value="unico">Único</option>
              <option value="gemelar">Gemelar</option>
              <option value="trillizos">Trillizos</option>
              <option value="cuatrillizos">Cuatrillizos</option>
            </b-select>
          </b-field>

          <b-field label="Puerpera aceptante">
            <b-select placeholder="Puerpera acentante" expanded>
              <option value="diu">DIU</option>
              <option value="otb">OTB</option>
              <option value="implante">Implante</option>
              <option value="bimensual">H.I. Bimensual</option>
              <option value="ningunoaceptante">Ninguno</option>
            </b-select>
          </b-field>

          <b-field label="Fecha de parto">
            <b-datepicker
              placeholder="Fecha del evento"
              icon="calendar-today"
              :locale="locale"
              disabled
              editable
            >
            </b-datepicker>
          </b-field>

          <b-field>
            <b-button @click="guardar()" type="is-success">Guardar</b-button>
          </b-field>
        </div>
      </div>
    </div>
  </b-steps>
</template>
<script>
import Utiles from "../services/Utiles";
import DialogosService from "../services/DialogosService";
import PuerperaService from "../services/PuerperaService";
export default {
  data: () => ({
    selected: null,
    edad: null,
    activeStep: 0,

    isProfileSuccess: false,
    checkbox: false,
    mobileMode: "minimalist",
    hasNavigation: true,
    locale: "es-MX",
    isAnimated: true,
    isRounded: true,
    isStepsClickable: false,
    prevIcon: "chevron-left",
    nextIcon: "chevron-right",
    showSocial: false,
    checkboxGroup: [],
    detalles: {
      APEO:"",
      consultasiete: null,
      consultaveinte: null,
      consultacuarenta: null,
    },
  }),
  computed: {
   
  },

  methods: {
    formatearFecha(fecha) {
      return Utiles.obtenerCadenaFecha(fecha);
    },
    async guardar() {
      if (!this.detalles.consultasiete == null) {
        return DialogosService.mostrarNotificacionError("Campos vacios");
      }
      const cargaUtil = {
        consultasiete: Utiles.obtenerCadenaFecha(this.detalles.consultasiete),
        consultaveinte:  Utiles.obtenerCadenaFecha(this.detalles.consultaveinte),
        consultacuarenta:  Utiles.obtenerCadenaFecha(this.detalles.consultacuarenta),
      };

      const respuesta = await PuerperaService.agregarPuerpera(cargaUtil);
      if (respuesta) {
        DialogosService.mostrarNotificacionExito("Vehículo registrado");
        this.detalles = {
          consultasiete: null,
          consultaveinte: null,
          consultacuarenta: null,
        };
      }
    },
  },
};
</script>