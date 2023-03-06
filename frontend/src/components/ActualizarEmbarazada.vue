<template>
  <!-- <div class="columns"> -->
  <b-steps v-model="activeStep" :animated="isAnimated" :rounded="isRounded" :has-navigation="hasNavigation"
    :icon-prev="prevIcon" :icon-next="nextIcon" :label-position="labelPosition" :mobile-mode="mobileMode">

    <b-step-item step="" label="" :clickable="isStepsClickable">
      <div class="column has-text-centered">
        <article class="panel centered">
          <p class="panel-heading">
            Actualizar datos de embarazada
          </p>
        </article>

      </div>
      <div>
        <div class="columns">
          <div class="column is-one-third">
            <b-field label="CURP">
              <b-input v-model="detalles.Curp" placeholder="Clave Unica de Registro de Población" icon="magnify"></b-input>
            </b-field>
            <b-field label="Nombre" :label-position="labelPosition">
              <b-input v-model="detalles.NombreEmbarazada" placeholder="Nombre completo de paciente"></b-input>
            </b-field>

            <b-field label="Número  teléfono">
              <b-input v-model="detalles.TelefonoEmbarazada" placeholder="Teléfono fijo"></b-input>
            </b-field>
            <b-field label="Municipio">
              <b-select placeholder="Municipio" expanded>
                <option value="Acatepec">Acatepec</option>
                <option value="Alcozauca">Alcozauca de Guerrero</option>
                <option value="Alpoyeca"> Alpoyeca</option>
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
                <option value="Otro">Otro</option>
              </b-select>
            </b-field>

            <b-field label="Comorbilidades">
              <section>
                <div class="block">
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
                  <b-checkbox v-model="checkboxGroup" native-value="Hepatitis">
                    Hepatitis
                  </b-checkbox>
                  <b-checkbox v-model="checkboxGroup" native-value="Ninguno">
                    Ninguno
                  </b-checkbox>
                </div>
                <p class="content">
                  <b>Selección:</b>
                  {{ checkboxGroup }}
                </p>
              </section>
            </b-field>
            <b-field label="Gestas">
              <b-input v-model="detalles.Gestas" placeholder="Número de gestas"></b-input>
            </b-field>
            <b-field label="Cesareas">
              <b-input v-model="detalles.Cesareas" placeholder="Número de cesareas"></b-input>
            </b-field>
          </div>

          <div class="column is-one-third">
            <b-field label="No. Progreso" :label-position="labelPosition">
              <b-input disabled v-model="detalles.id"></b-input>
            </b-field>
            <b-field label="Fecha de nacimiento">
              <b-datepicker placeholder="Fecha de nacimiento" icon="calendar-today" :locale="locale" editable>
              </b-datepicker>
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
            <b-field label="Localidad">
              <b-select placeholder="Localidad" expanded>
                <option value="Insabi">I</option>
              </b-select>
            </b-field>

            <b-field label="Detenciones"></b-field>
            <div class="field">
              <div class="b-checkbox">
                <input id="MamaNormal" class="styled" checked type="checkbox" />
                <label for="MamaNormal">
                  Exploración clínica de mama normal
                </label>
              </div>
              <div class="b-checkbox">
                <input id="MamaAlterado" class="styled" checked type="checkbox" />
                <label for="MamaAlterado">
                  Exploración clínica de mama alterado
                </label>
                <div class="b-checkbox">
                  <input id="PapaNormal" class="styled" checked type="checkbox" />
                  <label for="PapaNormal"> Papanicolau normal </label>
                </div>
                <div class="b-checkbox">
                  <input id="PapaLesion" class="styled" checked type="checkbox" />
                  <label for="PapaLesion"> Papanicolau con lesiones </label>
                </div>
                <div class="b-checkbox">
                  <input id="PapaCancer" class="styled" checked type="checkbox" />
                  <label for="PapaCancer"> Papanicolau cancer IN SITU </label>
                </div>

                <div class="b-checkbox">
                  <input id="PapaInvasor" class="styled" checked type="checkbox" />
                  <label for="PapaInvasor"> Papanicolau invasor </label>
                </div>
                <div class="b-checkbox">
                  <input id="NoDete" class="styled" checked type="checkbox" />
                  <label for="NoDete"> No se realizo </label>
                </div>
              </div>

              <b-field label="Paras">
                <b-input v-model="detalles.Paras" placeholder="Numero de paras"></b-input>
              </b-field>
              <b-field label="Abortos">
                <b-input v-model="detalles.Abortos" placeholder="Numero de abortos"></b-input>
              </b-field>
            </div>

          </div>
          <div class="column is-one-fifthr">
            <b-field label="No. Expediente clínico">
              <b-input v-model="detalles.Expediente" placeholder="No. Expediente clínico"></b-input>
            </b-field>
            <b-field label="Edad" :label-position="labelPosition">
              <b-input disabled="»disabled»" v-model="detalles.Edad"></b-input>
            </b-field>
            <div class="control">
              <b-field label="¿Emigro?"></b-field>
              <label class="radio">
                <input type="radio" name="Emigro" />
                Si
              </label>
              <label class="radio">
                <input type="radio" name="Emigro" />
                No
              </label>
              <b-input v-model="detalles.DondeMigro" placeholder="¿Dónde emigró?"></b-input>
            </div>
            <b-field label="Dirección">
              <b-input v-model="detalles.Direccion" placeholder="Dirección con referencia"></b-input>
            </b-field>
            <div class="control">
              <b-field label="¿Asistió a consulta pregestacional?"></b-field>
              <label class="radio">
                <input type="radio" name="ConsultaPreg" />
                Si
              </label>
              <label class="radio">
                <input type="radio" name="ConsultaPreg" />
                No
              </label>
            </div>

            <b-datepicker placeholder="Fecha de consulta" icon="calendar-today" :locale="locale" editable>
            </b-datepicker>


            <b-field label="Fecha de ultimo evento( Parto, cesarea, aborto)">
              <b-datepicker placeholder="Fecha de ultimo evento" icon="calendar-today" :locale="locale" editable>
              </b-datepicker>
            </b-field>

            <div class="control">
              <b-field label="¿Presentó complicaciones en el embarazo previo?"></b-field>
              <label class="radio">
                <input type="radio" name="Complicaciones" />
                Si
              </label>
              <label class="radio">
                <input type="radio" name="Complicaciones" />
                No
              </label>
              <b-select placeholder="¿Qué complicaciones presento?" expanded>
                <option value="Hemorragia ">Hemorragia </option>
                <option value="RetencioPlacenta">Retención de placenta</option>
                <option value="PlacentaPrevia">Placenta previa</option>
                <option value="DiabetesGesta">Diabetes gestacional</option>
                <option value="Preeclampsia">Preeclampsia</option>
                <option value="Eclampsia">Eclampsia </option>
                <option value="CovidEmbarazo">Covid + embarazo</option>
                <option value="OtraCompli">Otra</option>
                <option value="NingunaCompli">Ninguna</option>
              </b-select>
            </div>

          </div>
        </div>
      </div>
    </b-step-item>


    <b-step-item step="" label="" :clickable="isStepsClickable" :type="{ 'is-success': isProfileSuccess }">
      <div>
        <div class="columns">
          <div class="column is-one-third">

            <b-field label="Fecha de ultima menstruación">
              <b-datepicker placeholder="Fecha de ultima menstruación" icon="calendar-today" :locale="locale" editable>
              </b-datepicker>
            </b-field>


            <div class="control">
              <b-field label="Fecha de consutal">
                <b-datepicker placeholder="Fecha de consulta " icon="calendar-today" :locale="locale" editable>
                </b-datepicker>
              </b-field>

              <b-field label="Rubeola">
                <label class="radio">
                  <input type="radio" name="Rubeola" />
                  Si se aplicó previo al embarazo
                </label>
                <label class="radio">
                  <input type="radio" name="Rubeola" />
                  No se cuenta con vacuna
                </label>
              </b-field>


              <b-field label="Fecha de vacunación TD refuerzo">
                <b-datepicker placeholder="Fecha de vacunación TD refuerzo" icon="calendar-today" :locale="locale"
                  editable>
                </b-datepicker>

              </b-field>


              <b-field label="Vacuna COVID-19">
                <label class="radio">
                  <input type="radio" name="Covid" />
                  Se aplico
                </label>
                <label class="radio">
                  <input type="radio" name="Covid" />
                  No se aplico
                </label>
              </b-field>
              <b-field label="Biometria ematica">
                <b-select placeholder="Biometria" expanded>
                  <option value="HB12">HB >=12</option>
                  <option value="HBMAY9">HB >=9</option>
                  <option value="HBMEN9">HB Menor a 9 </option>
                  <option value="Ninguno">Ninguno</option>
                </b-select>
              </b-field>

              <b-field label="VIH">
                <b-select placeholder="VIH" expanded>
                  <option value="ReactivoVIH">Reactivo</option>
                  <option value="NoReactivoVIH">No reactivo</option>
                  <option value="NoRealizoVIH">No se realizó</option>
                </b-select>
              </b-field>
              <b-field label="Características fetales ">
                <b-select placeholder="Características" expanded>
                  <option value="UnicoVivo">Unico vivo</option>
                  <option value="UnicObito">Unico obito</option>
                  <option value="Molar">Molar</option>
                  <option value="Monocorial">Monocorial</option>
                  <option value="Bicorionicos">Bicorionicos</option>
                  <option value="Gemelar1Obitado">Gmelar 1 obitado</option>
                </b-select>
              </b-field>
            </div>
          </div>

          <div class="column is-one-third">
            <div class="control">
              <b-field label="Fecha probable de parto">
                <b-datepicker placeholder="Fecha probable de parto " icon="calendar-today" :locale="locale" editable>
                </b-datepicker>
              </b-field>
            </div>
            <b-field label="No. Total de consulta durante el embarazo" :label-position="labelPosition">
              <b-input disabled=»disabled» v-model="detalles.placas"
                placeholder="No. Total de consulta durante el embarazo"></b-input>
            </b-field>
            <div class="control">
              <b-field label="Fecha de vacunación TD primera">
                <b-datepicker placeholder="Fecha de vacunación TD primera " icon="calendar-today" :locale="locale"
                  editable>
                </b-datepicker>
              </b-field>
            </div>
            <div class="control">
              <b-field label="Fecha de vacunación TDPA">
                <b-datepicker placeholder="Fecha de vacunación TDPA " icon="calendar-today" :locale="locale" editable>
                </b-datepicker>
              </b-field>
            </div>

            <b-field label="Grupo RH">
              <b-select placeholder="Grupo RH" expanded>
                <option value="A+">A+</option>
                <option value="A-">A-</option>
                <option value="B+">B+</option>
                <option value="B-">B-</option>
                <option value="O+">O+</option>
                <option value="O-">O-</option>
                <option value="AB+">AB+</option>
                <option value="AB-">AB-</option>
                <option value="NingunoRH">Ninguno</option>
              </b-select>
            </b-field>
            <b-field label="Leucocitos">
              <b-select placeholder="Leucocitos" expanded>
                <option value="Normales">Normales</option>
                <option value="DisminuidosLinfocitos">Disminuidos linfocitos</option>
                <option value="AumentadosNeutrofilos">Aumentados neutrofilos</option>
                <option value="AumentadosLinfocitos">Aumentados los linfocitos</option>
                <option value="NingunoLEUCO">Ninguno</option>
              </b-select>
            </b-field>

            <b-field label="Estado tiras glucosa">
              <b-select placeholder="Estado en el que se realizó" expanded>
                <option value="Ayuno">Ayuno</option>
                <option value="Posprandial1">Posprandial 1 hora</option>
                <option value="Posprandial2">Posprandial 2 horas</option>
                <option value="NingunoLEUCO">Ninguno</option>
              </b-select>
            </b-field>

            <b-field label=" Malformaciones congenitas">
              <b-select placeholder="Malformaciones" expanded>
                <option value="EspinaBifida">Espina bifida</option>
                <option value="Meningocele">Meningocele</option>
                <option value="Mielomeningocele">Mielomeningocele</option>
                <option value="Anencefalia">Anencefalia</option>
                <option value="OtrasMalform">Otras</option>
                <option value="NingunaMalform">Ninguna</option>
              </b-select>
            </b-field>

          </div>
          <div class="column is-one-third">

            <b-field label="Semana de gestación actual" :label-position="labelPosition">
              <b-input disabled=»disabled» v-model="detalles.placas" placeholder="Semanan de gestación actual">
              </b-input>
            </b-field>
            <b-field label="No. Total de consulta en el mes" :label-position="labelPosition">
              <b-input disabled=»disabled» v-model="detalles.placas" placeholder="No. Total de consulta en el mes ">
              </b-input>
            </b-field>
            <div class="control">
              <b-field label="Fecha de vacunación TD segunda">
                <b-datepicker placeholder="Fecha de vacunación TD segunda " icon="calendar-today" :locale="locale"
                  editable>
                </b-datepicker>
              </b-field>
            </div>
            <div class="control">
              <b-field label="Fecha de vacunación Influenza estacional">
                <b-datepicker placeholder="Fecha de vacunación Influenza estacional " icon="calendar-today"
                  :locale="locale" editable>
                </b-datepicker>
              </b-field>
            </div>
            <b-field label="EGO y/o tiras urianalisis">
              <b-select placeholder="EGO" expanded>
                <option value="Trazas">Prot. trazas</option>
                <option value="Mayor300">Prot. >300 mg/l</option>
                <option value="Leucitos">leucitos >5</option>
                <option value="Nitratos">nitratos +</option>
                <option value="NegativoEgo">Negativo</option>
              </b-select>
            </b-field>

            <b-field label="VDRL">
              <b-select placeholder="VDRL" expanded>
                <option value="ReactivoVDRL">Reactivo</option>
                <option value="NoReactivoVDRL">No reactivo</option>
                <option value="NoRealizo">No se realizo</option>
              </b-select>
            </b-field>
            <b-field label="Resultado tiras glucosa">
              <b-select placeholder="Resultado tiras glucosa" expanded>
                <option value="Mayor60">>60 y menor a 110</option>
                <option value="Mayor110">>=110 y menor o igual a 126</option>
                <option value="Mayor126">>126 y menor a 160</option>
                <option value="NingunoGlucosa">Ninguno</option>
              </b-select>
            </b-field>

            <b-field label="Liquido amniotico">
              <b-select placeholder="Liquido amniotico" expanded>
                <option value="Normal">Normal</option>
                <option value="Oligohidramnios">Oligohidramnios</option>
                <option value="Polihidramnios">Polihidramnios</option>
                <option value="Anhidramnios">Anhidramnios</option>
              </b-select>
            </b-field>
          </div>
        </div>
      </div>
    </b-step-item>

    <b-step-item step="" label="" :clickable="isStepsClickable" :type="{ 'is-success': isProfileSuccess }">
      <div>
        <div class="columns">
          <div class="column is-one-third">
            <b-field label="Placenta">
              <b-select placeholder="Placenta" expanded>
                <option value="NormalPlacen">Normal</option>
                <option value="Previa">Previa</option>
                <option value="AcretismoPlacentario">Acretismo placentario</option>
                <option value="Calcificada">Calcificada</option>
              </b-select>
            </b-field>


            <b-field label="¿Acudió a referencia?"></b-field>
            <b-select placeholder="Referencias" expanded>
              <option value="HgTlapa ">H.G. Tlapa</option>
              <option value="Hmnig">HMNIG</option>
              <option value="HcAcatepec">H.C. Acatepec</option>
              <option value="HcAlcozauca">H.C. Alcozauca</option>
              <option value="HcMalina">H.C. Malinaltepec </option>
              <option value="HcOlinala">H.C. Olinalá</option>
              <option value="HcTlacoapa">H.C. Tlacoapa</option>
              <option value="HcXochi">H.C. Xochihuehuetlán</option>
              <option value="HcZapo">H.C. Zapotitlán tablas</option>
              <option value="CentroSalud">Centro de salud</option>
              <option value="Hogara">HOGAR</option>
              <option value="HcSaLuis">H.C. San sluis</option>
              <option value="HgAyutla">H.G. Ayutla</option>
              <option value="HgOme">H.G. Ometepec</option>
              <option value="HgChuautla">H.G Chuautla</option>
              <option value="OtrosHospitales">Otros hospitales</option>
              <option value="NoAcudio">No acudió</option>
            </b-select>


            <b-field label="¿Se oriento en señales de peligro durante el embarazo, parto, puerperio, rn?"></b-field>
            <b-select placeholder="¿Se oriento en señales de peligro durante el embarazo?" expanded>
              <option value="Hemorragia ">Si</option>
              <option value="Retención de placenta">No</option>
            </b-select>
            <b-field label="¿cuenta con transporte ame en caso de emergencia?">
              <b-select placeholder="¿Cuenta con transporte AME?" expanded>
                <option value="SiAME">SI</option>
                <option value="NoAME">NO</option>
              </b-select>
            </b-field>
          </div>

          <div class="column is-one-third">
            <div class="control">
              <b-field label="Fecha probable de parto por USG">
                <b-datepicker placeholder="Fecha probable de parto por USG (1Er Trimestre)" icon="calendar-today"
                  :locale="locale">
                </b-datepicker>
              </b-field>
            </div>
            <b-field label="Cntrareferencia recibida">
              <b-select placeholder="Referencia" expanded>
                <option value="SiConrefe">Si</option>
                <option value="NoCorefe">No</option>
              </b-select>
            </b-field>
            <b-field label="¿Donde atenderán su parto?">
              <b-select placeholder="¿Donde atenderán su parto?" expanded>
                <option value="AtenderHospital">Hospital</option>
                <option value="AtenderC.Salud">C. Salud</option>
                <option value="AtenderHogar">Hogar</option>
                <option value="AtenderParticular">Particula
                </option>
              </b-select>
            </b-field>
            <div class="control">
              <b-field label="Fecha de evento">
                <b-datepicker placeholder="Fecha de evento" icon="calendar-today" :locale="locale" editable>
                </b-datepicker>
              </b-field>
            </div>

          </div>
          <div class="column is-one-third">
            <b-field label="Referencia">
              <b-select placeholder="Referencia" expanded>
                <option value="SiRefe">Si</option>
                <option value="NoRefe">No</option>
              </b-select>
            </b-field>

            <b-field label="Se le realizó el plan de seguridad">
              <b-select placeholder="Plan de seguridad" expanded>
                <option value="SiRefe">Si</option>
                <option value="NoRefe">Reforzamiento</option>
                <option value="NoRefe">No</option>
              </b-select>
            </b-field>

            <b-field label="¿Quien atenderá su parto?">
              <b-select placeholder="¿Quien atenderá su parto?" expanded>
                <option value="AtenderMedico">Medico</option>
                <option value="AtenderEnfra">Enfra</option>
                <option value="AtenderParteraProfs">Partera profes.</option>
                <option value="AtenderParteraTradc">Partera tradic.</option>
                <option value="AtenderFamiliar">Familiar</option>
                <option value="AtenderOtro">Otro</option>
              </b-select>
            </b-field>

          </div>
        </div>
      </div>
    </b-step-item>
    <b-step-item :step="showSocial ? '' : ''" label="" :clickable="isStepsClickable" disabled>
      <div class="column has-text-centered">
        <h3 class="is-size-3">Triagge</h3>
        <p>Registrar embarazada</p>
      </div>
      <table class="table is-bordered is-striped is-narrow is-hoverable is-fullwidth">
        <thead>
          <tr>
            <th class="is-warning"><abbr title="DatosGene"></abbr>DATOS GENERALES</th>
            <th class="is-warning"> <abbr title="BajoRiesgo"></abbr>BAJO RIESGO</th>
            <th class="is-warning"><abbr title="AltoRiesgo" center></abbr>ALTO RIESGO</th>
            <th class="is-warning"><abbr title="AltoRiesgo" center></abbr></th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Edad</td>
            <td class="is-success"> <b-field>
                <b-checkbox v-model="checkbox" native-value="EdadBajo" type="is-success">21 a 34 años</b-checkbox>
              </b-field></td>
            <td><b-field>
                <b-checkbox v-model="checkbox" native-value="EdadAlto" type="is-danger">Menor a 17 o mayor a 35
                  años</b-checkbox>
              </b-field></td>
            <td></td>
          </tr>
          <tr>
            <td>IMC al inicio del embarazo</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ImcBajo" type="is-success">18.5 a 24.9 P.
                  ideal</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ImcAlto" type="is-danger">Menor a 17 o Mayor a
                  30Kg/E</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Talla</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="TallaBajo" type="is-success">Menor o igual a
                  145cm</b-checkbox>
              </b-field></td>
            <td><b-field>
                <b-checkbox v-model="checkbox" native-value="TallaAlto" type="is-danger">>145</b-checkbox>
              </b-field></td>
            <td></td>
          </tr>

          <tr>
            <td>Esacolarida</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="EscolaridadBajo" type="is-success">Mayor a
                  Primaria</b-checkbox>
              </b-field></td>
            <td><b-field>
                <b-checkbox v-model="checkbox" native-value="EscolaridadAlto" type="is-danger">Nula</b-checkbox>
              </b-field></td>
            <td></td>
          </tr>
          <tr>
            <td>Trabajar más de 36hrs/10hrs al día de pie o sentada</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="TrabajarBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="TrabajarAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Exposición a tóxicos(plomo, pesticidas, quimicos, Enf.Infecciosas)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <th class="is-warning"><abbr title="Antecedentes"></abbr>ANTECEDENTES</th>
            <th class="is-warning"> <abbr title="BajoRiesgo"></abbr>BAJO RIESGO</th>
            <th class="is-warning"><abbr title="AltoRiesgo" center></abbr>ALTO RIESGO</th>
            <th class="is-warning"><abbr title="AltoRiesgo" center></abbr></th>
          </tr>
          <tr>
            <td>Aborto previos Mayo o igual a 3(LUI)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="AbortoBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="AbortoAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Partos prematuro (Menor a 28 SDG) previos menor o igual a 1</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="PartosBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="PartosAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Muertes perinatales u Óbitos</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="MuertesBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="MuertesAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Hijos con malformaciones</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="MarlformacionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="MarlformacionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Infección de vías urinarias(IVU Refactorio o recurrente)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">IVU baja</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Piolonefritis</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Periodo intergenésico corto</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Mayor o igual a 2
                  años</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Menor a 12
                  meses</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Paridad</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">1 a 3</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Mayor o igual
                  a4</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Cesárea anterior</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Mayor a 1</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Parto anterior complicado (Hemorragia obstétrica, Retencion placentaria, TDPP)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Sin
                  complicaciones</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Complicado</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Preclampsia-eclampsia</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Productos previos</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Peso normal</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Menor a
                  2.5kg(restricción)/Mayor a 4hkg(macrosomico)</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Cirugía genitourinaria(Miomectomía, conizacíon, Ectópico roto, etc.)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Comorbilidades</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Ninguna</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <th class="is-warning"><abbr title="DatosGene"></abbr>EMBARAZO ACTUAL</th>
            <th class="is-warning"> <abbr title="BajoRiesgo"></abbr>BAJO RIESGO</th>
            <th class="is-warning"><abbr title="AltoRiesgo" center></abbr>ALTO RIESGO</th>
            <th class="is-warning"><abbr title="AltoRiesgo" center></abbr></th>
          </tr>
          <tr>
            <td>Toxicomanias(Cocaína, tabaco, alcohol, crack, etc)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Hepatitis B y C</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>"Rh (·} Profilaxis prenatal rutinaria con una administración de 300 µg (1.500 UI) de
              inmunoalobulina anti-O, a la 24-30 SDG, 72 hrs. Posterior Evento Obstétrico"
            </td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>VIH(+)/SIFILIS(+)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Hemoglobina en gr. 1er Trimestre del embarazo.</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">>11.5
                </b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Menor a 9
                </b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Retardo menstrual +Dolor Abdominal (Descartar Embarazo Ectópico roto)
            </td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-danger"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Hiperémesis gravídica</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No presente</b-checkbox>
              </b-field></td>
            <td class="is-danger"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Presente</b-checkbox>
              </b-field></td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Fondo Uterino/ SDG (Normal: 7-13 cm 1er;20- 26.8 cm 2do y 30-39 cm 3er Trimestre) Descartar:
              Restricción; prematuro 6 Macrosomico; polidramnlos; Embarazo Molar 6 Múltiple
            </td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Dentro Percentil (10 -
                  90)
                </b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Por debajo de percentil 10
                  o por arriba del 90</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>"Descartar Diabetes gestacional a las 24-28 SDG (> 126 mg/dl ayuno, CTG 50gr.
              >200 ma/dl al Azar; > 200 ma/dl 2 horas (75 ar CTG)"
            </td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Mujeres con IMC >30 al inicio del embarazo, registre una ganancia de peso</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">4 A 9 Kg</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">> 9 Kg (IMC> 31.7 2DO Y >
                  33 3ER TRIMESTRE
                </b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Amenaza de aborto/Menor a 500 gr o ≥ 22 SDG</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Infección Urinaria o Vaginosis</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Refactorio o
                  recurrente</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>"Amenaza Parto Prematuro (22-36.6 SDG, contracciones, dilatación ≥3 cm._borramiento
              90%, con o sin membranas amnióticas integras)"
            </td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Enfermedad Hipertensiva del embarazo/ Preclamsia</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">TA 100/60 -
                  120/70</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">130/80 - 140 - 90
                </b-checkbox></b-field>
            </td>
            <td class="is-danger">
              <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">160-110 MMHG
              </b-checkbox>
            </td>
          </tr>
          <tr>
            <td>Proteinuria de Orina</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Normal</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">+ A + + +</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Datos de severidad (cefalea, nauseas. dolor en epigastrio, edema, etc.)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Ausentes</b-checkbox>
              </b-field></td>
            <td class="is-danger"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Presentes</b-checkbox>
              </b-field></td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Crisis Convulsivas</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Ausentes</b-checkbox>
              </b-field></td>
            <td class="is-danger"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Presentes</b-checkbox>
              </b-field></td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Malformaciones Conqénitas</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Determinar posición fetal</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Cefálico</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Pélvico, Transverso u
                  oblicuo</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Movimientos fetales activos(+ 2 horas. después 28 SDG)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Frecuencia Cardiaca Fetal</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">120 A 160 LAT X
                  MIN</b-checkbox>
              </b-field></td>
            <td class="is-danger"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Menor a 100 O >160 LATIDOS
                  X MIN</b-checkbox>
              </b-field></td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Implantación anómala de la Placenta</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si (Centarl al orificio
                  cervical, marginal, acreta, etc.)</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Hemorragia transvaginal (Obstétrica)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-danger"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si (Descartar DPPNI,
                  Placenta previa)</b-checkbox>
              </b-field></td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Taquicardia</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">100 a 119</b-checkbox>
              </b-field>
            </td>
            <td class="is-danger">
              <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Mayor o igual a 120 latidos
                X min</b-checkbox>
            </td>
          </tr>
          <tr>
            <td>Frecuencia Respiratoria</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">11 A 20 X
                  MINUTO</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">21 A 30</b-checkbox>
              </b-field></td>
            <td class="is-danger">
              <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">≤ 10 Ó ≥ 30 X
                MIN</b-checkbox>
            </td>
          </tr>
          <tr>
            <td>Respuesta neurológica</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Alerta</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Voz</b-checkbox>
              </b-field></td>
            <td class="is-danger"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Responde al
                  dolor</b-checkbox>
              </b-field></td>
          </tr>
          <tr>
            <td>Ruptura prematura de membranas</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Menor a 18
                  HORAS</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">≥ 18 HORAS</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Embarazo pos-termino 41 SDG sin trabajo de parto</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Fiebre > 38 ºC</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Síntomas Respiratorios (Tos seca, cefalea, anosmia, rinorrea, faringitis, etc.)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">No</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Si</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Concentración de Oxígeno /Dificultad para respirar, descartar Covid-19)</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">90 a 100%</b-checkbox>
              </b-field></td>
            <td class="is-danger"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Menos de 90</b-checkbox>
              </b-field></td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Redes de apoyo</td>
            <td class="is-success"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionBajo" type="is-success">Presente</b-checkbox>
              </b-field></td>
            <td class="is-warning"><b-field>
                <b-checkbox v-model="checkbox" native-value="ExposicionAlto" type="is-danger">Ausente</b-checkbox>
              </b-field></td>
            <td class="is-warning"></td>
          </tr>
        </tbody>
      </table>
      <b-field>
        <b-button @click="guardar()" type="is-success">Guardar</b-button>
        <router-link :to="{ name: 'Vehiculos' }" class="button is-info ml-2">
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
import EmbarazadaService from "../services/EmbarazadaService";
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
      const respuesta = await EmbarazadaService.agregarVehiculo(cargaUtil);
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