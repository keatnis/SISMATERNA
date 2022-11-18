<template>
  <!-- <div class="columns"> -->
  <b-steps v-model="activeStep" :animated="isAnimated" :rounded="isRounded" :has-navigation="hasNavigation"
    :icon-prev="prevIcon" :icon-next="nextIcon" :label-position="labelPosition" :mobile-mode="mobileMode">

    <b-step-item step="1" label="Datos Generales" :clickable="isStepsClickable">
      <article class="panel is-primary centered">
        <p class="panel-heading">
          Actualizar embarazada
        </p>
      </article>
        <div class="column has-text-centered">
          <h3 class="is-size-3">Datos generales</h3>
        </div>

        <b-field label="CURP"></b-field>
        <div class="control">
          <p class="control has-icons-left">
            <input class="input" type="text" placeholder="Buscar por CURP">
            <span class="icon is-left">
              <i class="fas fa-search" aria-hidden="true"></i>
            </span>
          </p>
        </div>
        <div class="control">
          <b-field label="No. Progreso" :label-position="labelPosition">
            <b-input disabled="»disabled»" v-model="detalles.placas" placeholder=""></b-input>
          </b-field>

        </div>

        <b-field label="No. Expediente clínico">
          <b-input v-model="detalles.descripcion" placeholder="No. Expediente clínico"></b-input>
        </b-field>

        <b-field label="Nombre" :label-position="labelPosition">
          <b-input v-model="detalles.placas" placeholder="Nombre completo de paciente"></b-input>
        </b-field>

        <b-field label="Fecha de nacimiento">
          <b-datepicker placeholder="Fecha de nacimiento" icon="calendar-today" :locale="locale" editable>
          </b-datepicker>
        </b-field>

        <b-field label="Edad" :label-position="labelPosition">
          <b-input disabled="»disabled»" v-model="detalles.placas"></b-input>
        </b-field>

        <b-field label="Número  teléfono">
          <b-input v-model="detalles.descripcion" placeholder="Teléfono fijo"></b-input>
        </b-field>

        <b-field label="Lengua Indígena">
          <b-select placeholder="Lengua Indígena" expanded>
            <option value="Nahual">Nahual</option>
            <option value="Mixteco">Tu´un Savi (Mixteco)</option>
            <option value="Mephaa">Meephaa (Tlapaneco)</option>
            <option value="Amuzgo">Amuzgo</option>
            <option value="Ninguno">Ninguno</option>
          </b-select>
        </b-field>

        <b-field label="Derechohabiencia">
          <b-select placeholder="Derechohabiencia" expanded>
            <option value="Insabi">INSABI</option>
            <option value="Imss">IMSS</option>
            <option value="Issste">ISSSTE</option>
            <option value="Sedena">SEDENA</option>
            <option value="Pemex">PEMEX</option>
            <option value="Marina">MARINA</option>
            <option value="Otro">OTRO</option>
          </b-select>
        </b-field>

        <div class="control">
          <b-field label="¿Emigro?"></b-field>
          <label class="radio">
            <input type="radio" name="answer" />
            Si
          </label>
          <label class="radio">
            <input type="radio" name="answer" />
            No
          </label>
          <b-input v-model="detalles.descripcion" placeholder="¿Dónde emigró?"></b-input>
        </div>

        <h2 class="subtitle"></h2>
        <h2 class="subtitle">DOMICILIO</h2>
        <b-field label="Dirección">
          <b-input v-model="detalles.descripcion" placeholder="Dirección con referencia"></b-input>
        </b-field>
        <b-field label="Localidad">
          <b-input v-model="detalles.descripcion" placeholder="Localidad"></b-input>
        </b-field>
        <div></div>
        <b-field label="Municipio">
          <b-input v-model="detalles.descripcion" placeholder="Municipio"></b-input>
        </b-field>
        <div></div>
        <b-field label="Detenciones"></b-field>
        <div class="field">
          <div class="b-checkbox">
            <input id="checkbox1" class="styled" checked type="checkbox" />
            <label for="checkbox1">
              Exploración clínica de mama normal
            </label>
          </div>
          <div class="b-checkbox">
            <input id="checkbox2" class="styled" checked type="checkbox" />
            <label for="checkbox2">
              Exploración clínica de mama alterado
            </label>
            <div class="b-checkbox">
              <input id="checkbox1" class="styled" checked type="checkbox" />
              <label for="checkbox1"> Papanicolau normal </label>
            </div>
            <div class="b-checkbox">
              <input id="checkbox1" class="styled" checked type="checkbox" />
              <label for="checkbox1"> Papanicolau con lesiones </label>
            </div>
            <div class="b-checkbox">
              <input id="checkbox1" class="styled" checked type="checkbox" />
              <label for="checkbox1"> Papanicolau cancer IN SITU </label>
            </div>

            <div class="b-checkbox">
              <input id="checkbox1" class="styled" checked type="checkbox" />
              <label for="checkbox1"> Papanicolau invasor </label>
            </div>
            <div class="b-checkbox">
              <input id="checkbox1" class="styled" checked type="checkbox" />
              <label for="checkbox1"> No se realizo </label>
            </div>
          </div>

          <div class="column"></div>
          <div class="control">
            <b-field label="¿Asistió a consulta pregestacional?"></b-field>
            <label class="radio">
              <input type="radio" name="answer" />
              Si
            </label>
            <label class="radio">
              <input type="radio" name="answer" />
              No
            </label>
          </div>
          <b-field label="">
            <b-datepicker placeholder="Fecha de consulta" icon="calendar-today" :locale="locale" editable>
            </b-datepicker>
          </b-field>
        </div>
      
      <section>
        <div class="block">
          <h2 class="subtitle">Comorbilidades </h2>
          <b-checkbox v-model="checkboxGroup" native-value="sobrepeso">
            Sobre peso
          </b-checkbox>
          <b-checkbox v-model="checkboxGroup" native-value="Obesidad">
            Obesidad
          </b-checkbox>
          <b-checkbox v-model="checkboxGroup" native-value="Diabetes Mellitus">
            Diabetes Mellitus
          </b-checkbox>
          <b-checkbox v-model="checkboxGroup" native-value="Hipertensión arterial">
            Hipertensión arterial
          </b-checkbox>
          <b-checkbox v-model="checkboxGroup" native-value="Cardiopatía">
            Cardiopatía
          </b-checkbox>
          <b-checkbox v-model="checkboxGroup" native-value="Epilepsia ">
            Epilepsia
          </b-checkbox>
          <b-checkbox v-model="checkboxGroup" native-value="Sifilis ">
            Sifilis
          </b-checkbox>
          <b-checkbox v-model="checkboxGroup" native-value="VIH ">
            VIH
          </b-checkbox>
          <b-checkbox v-model="checkboxGroup" native-value="eHepatitis">
            Hepatitis
          </b-checkbox>
        </div>
        <p class="content">
          <b>Selección:</b>
          {{ checkboxGroup }}
        </p>
      </section>
    </b-step-item>

    <b-step-item step="2" label="Antecedentes" :clickable="isStepsClickable" :type="{ 'is-success': isProfileSuccess }">
      <div class="column has-text-centered">
        <h3 class="is-size-3">Antecedentes gineco obstétricos</h3>
        <p>Actualizar embarazada</p>
      </div>

      <div class="columns">
        <div class="column">

          <b-field label="Gestas">
            <b-input v-model="detalles.descripcion" placeholder="Numero de gestas"></b-input>
          </b-field>

          <b-field label="Paras">
            <b-input v-model="detalles.descripcion" placeholder="Numero de paras"></b-input>
          </b-field>
        </div>
      </div>
      <b-field label="Abortos">
        <b-input v-model="detalles.descripcion" placeholder="Numero de abortos"></b-input>
      </b-field>
      <b-field label="Cesareas">
        <b-input v-model="detalles.descripcion" placeholder="Numero de cesareas"></b-input>
      </b-field>
      <div>
        <b-field label="Fecha de ultimo evento( Parto, cesarea, aborto)">
          <b-datepicker placeholder="Fecha de Nacimiento" icon="calendar-today" :locale="locale" editable>
          </b-datepicker>
        </b-field>
      </div>
      <div class="control">
        <b-field label="¿Presentó complicaciones en el embarazo previo?"></b-field>
        <label class="radio">
          <input type="radio" name="answer" />
          Si
        </label>
        <label class="radio">
          <input type="radio" name="answer" />
          No
        </label>
        <b-select placeholder="¿Qué complicaciones presento?" expanded>
          <option value="Hemorragia ">Hemorragia </option>
          <option value="Retención de placenta">Retención de placenta</option>
          <option value="Placenta previa">Placenta previa</option>
          <option value="Diabetes gestacional">Diabetes gestacional</option>
          <option value="Preeclampsia">Preeclampsia</option>
          <option value="Preeclampsia">Eclampsia </option>
          <option value="Preeclampsia">Covid + embarazo</option>
          <option value="Preeclampsia">Otra</option>
          <option value="Preeclampsia">Ninguno</option>
        </b-select>
      </div>
    </b-step-item>

    <b-step-item step="3" label="Embarazo actual" :clickable="isStepsClickable"
      :type="{ 'is-success': isProfileSuccess }">
      <div class="column has-text-centered">
        <h3 class="is-size-3">Embarazo actual</h3>
        <p>Actualizar embarazada</p>
      </div>

      <div class="column"> </div>

      <div class="control">
        <b-field label="Fecha de ultima menstruación">
          <b-datepicker placeholder="Fecha de ultima menstruación" icon="calendar-today" :locale="locale" editable>
          </b-datepicker>
        </b-field>
      </div>

      <div class="control">
        <b-field label="Fecha probable de parto">
          <b-datepicker placeholder="Fecha probable de parto " icon="calendar-today" :locale="locale" editable>
          </b-datepicker>
        </b-field>
      </div>

      <div>
        <b-field label="Semana de gestación actual" :label-position="labelPosition">
          <b-input disabled=»disabled» v-model="detalles.placas" placeholder="Semanan de gestación actual"></b-input>
        </b-field>
      </div>

      <div class="control">
        <b-field label="Fecha de consutal">
          <b-datepicker placeholder="Fecha de consulta " icon="calendar-today" :locale="locale" editable>
          </b-datepicker>
        </b-field>
      </div>
      <div>
        <b-field label="No. Total de consulta durante el embarazo" :label-position="labelPosition">
          <b-input disabled=»disabled» v-model="detalles.placas"
            placeholder="No. Total de consulta durante el embarazo"></b-input>
        </b-field>
      </div>
      <div>
        <b-field label="No. Total de consulta en el mes" :label-position="labelPosition">
          <b-input disabled=»disabled» v-model="detalles.placas" placeholder="No. Total de consulta en el mes ">
          </b-input>
        </b-field>
      </div>

    </b-step-item>

    <b-step-item step="4" :visible="showSocial" label="Esquema de vacunación" :clickable="isStepsClickable">
      <div class="column has-text-centered">
        <h3 class="is-size-3">Ezquema de vacunación</h3>
        <p>Actualizar embarazada</p>
      </div>

      <div class="control">
        <b-field label="Rubeola"></b-field>
        <label class="radio">
          <input type="radio" name="answer" />
          Si se aplicó previo al embarazo
        </label>
        <label class="radio">
          <input type="radio" name="answer" />
          No se cuenta con vacuna
        </label>
      </div>
      <div class="control">
        <b-field label="Fecha de vacunación TD primera">
          <b-datepicker placeholder="Fecha de vacunación TD primera " icon="calendar-today" :locale="locale" editable>
          </b-datepicker>
        </b-field>
      </div>

      <div class="control">
        <b-field label="Fecha de vacunación TD segunda">
          <b-datepicker placeholder="Fecha de vacunación TD segunda " icon="calendar-today" :locale="locale" editable>
          </b-datepicker>
        </b-field>
      </div>
      <div class="control">
        <b-field label="Fecha de vacunación TD refuerzo">
          <b-datepicker placeholder="Fecha de vacunación TD refuerzo" icon="calendar-today" :locale="locale" editable>
          </b-datepicker>
        </b-field>
      </div>
      <div class="control">
        <b-field label="Fecha de vacunación TDPA">
          <b-datepicker placeholder="Fecha de vacunación TDPA " icon="calendar-today" :locale="locale" editable>
          </b-datepicker>
        </b-field>
      </div>
      <div class="control">
        <b-field label="Fecha de vacunación Influenza estacional">
          <b-datepicker placeholder="Fecha de vacunación Influenza estacional " icon="calendar-today" :locale="locale"
            editable>
          </b-datepicker>
        </b-field>
      </div>

      <b-field label="Vacuna COVID-19"></b-field>
      <label class="radio">
        <input type="radio" name="answer" />
        Se aplico
      </label>
      <label class="radio">
        <input type="radio" name="answer" />
        No se aplico
      </label>
    </b-step-item>

    <b-step-item step="5" :visible="showSocial" label="Laboratorio" :clickable="isStepsClickable">
      <h1 class="title has-text-centered">Actualizar laboratorio</h1>
      <div class="column has-text-centered">
        <h3 class="is-size-3">Laboratorio</h3>
        <p>Actualizar embarazada</p>
      </div>
    </b-step-item>
    <b-step-item step="6" :visible="showSocial" label="Trabajo social" :clickable="isStepsClickable">
      <div class="column has-text-centered">
        <h3 class="is-size-3">Trabajo social</h3>
        <p>Actualizar embarazada</p>
      </div>
    </b-step-item>
    <b-step-item step="7" :visible="showSocial" label="Componente comunitario" :clickable="isStepsClickable">
      <div class="column has-text-centered">
        <h3 class="is-size-3">Componente comunitario</h3>
        <p>Actualizar embarazada</p>
      </div>
    </b-step-item>
    <b-step-item :step="showSocial ? '8' : '8'" label="Trigge" :clickable="isStepsClickable" disabled>
      <div class="column has-text-centered">
        <h3 class="is-size-3">Triagge</h3>
        <p>Actualizar embarazada</p>
      </div>

      <b-field>
        <b-button @click="guardar()" type="is-success">Guardar</b-button>
        <router-link :to="{ name: '1' }" class="button is-info ml-2">
          Volver
        </router-link>
      </b-field>
    </b-step-item>

    <template v-if="customNavigation" #navigation="{ previous, next }">
      <b-button outlined type="is-danger" icon-pack="fas" icon-left="backward" :disabled="previous.disabled"
        @click.prevent="previous.action">
        Previous
      </b-button>
      <b-button outlined type="is-success" icon-pack="fas" icon-right="forward" :disabled="next.disabled"
        @click.prevent="next.action">
        Next
      </b-button>
    </template>
  </b-steps>
</template>
<script>
//labelPosition: 'on-border',
import Utiles from "../services/Utiles";
import DialogosService from "../services/DialogosService";
import VehiculosService from "../services/VehiculosService";
export default {
  data: () => ({
    selected: null,
    fechaYHoraActual: null,
    usarFechaYHoraActual: true,
    checkboxGroup: [],
    detalles: {
      placas: "",
      descripcion: "",
      fechaEntrada: null,

    },
  }),
  computed: {
    selectedString() {
      return this.selected ? this.selected.toDateString() : "";
    },
  },
  mounted() {
    this.detalles.fechaEntrada = new Date();
    this.refrescarFechaYHoraActual();
    setInterval(this.refrescarFechaYHoraActual, 500);
  },
  methods: {
    refrescarFechaYHoraActual() {
      this.fechaYHoraActual = Utiles.obtenerFechaYHoraActual();
    },
    formatearFecha(fecha) {
      return Utiles.obtenerFechaYHora(fecha);
    },
    async guardar() {
      if (!this.usarFechaYHoraActual && !this.detalles.fechaEntrada) {
        return DialogosService.mostrarNotificacionError(
          "Selecciona una fecha y hora"
        );
      }
      const cargaUtil = {
        placas: this.detalles.placas,
        descripcion: this.detalles.descripcion,
        fechaEntrada: Utiles.obtenerFechaYHora(this.detalles.fechaEntrada, "T"),
      };
      if (this.usarFechaYHoraActual) {
        cargaUtil.fechaEntrada = Utiles.obtenerFechaYHoraActual("T");
      }
      const respuesta = await VehiculosService.agregarVehiculo(cargaUtil);
      if (respuesta) {
        DialogosService.mostrarNotificacionExito("Vehículo registrado");
        this.detalles = {
          placas: "",
          descripcion: "",
          fechaEntrada: new Date(),
        };
      }
    },
  },
};
</script>
<style>

</style>