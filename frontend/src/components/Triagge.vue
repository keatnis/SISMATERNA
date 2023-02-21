<template >
    <div>

     <div class="column has-text-centered">
        <h3 class="is-size-3">Triagge</h3>
        <p>Registrar embarazada</p>
      </div>
      <table
        class="table is-bordered is-striped is-narrow is-hoverable is-fullwidth"
      >
        <thead>
          <tr>
            <th class="is-warning">
              <abbr title="DatosGene"></abbr>DATOS GENERALES
            </th>
            <th class="is-warning">
              <abbr title="BajoRiesgo"></abbr>BAJO RIESGO
            </th>
            <th class="is-warning">
              <abbr title="AltoRiesgo" center></abbr>ALTO RIESGO
            </th>
            <th class="is-warning"><abbr title="AltoRiesgo" center></abbr></th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Edad</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="edadBajo"
                  native-value="EdadBajo"
                  type="is-success"
                  >21 a 34 años</b-checkbox
                >
              </b-field>
            </td>
            <td>
              <b-field>
                <b-checkbox
                  v-model="edadAlta"
                  native-value="EdadAlto"
                  type="is-danger"
                  >Menor a 17 o mayor a 35 años</b-checkbox
                >
              </b-field>
            </td>
            <td></td>
          </tr>
          <tr>
            <td>IMC al inicio del embarazo</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="ImcBajo"
                  native-value="ImcBajo"
                  type="is-success"
                  >18.5 a 24.9 P. ideal</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="ImcAlto"
                  native-value="ImcAlto"
                  type="is-danger"
                  >Menor a 17 o Mayor a 30Kg/E</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Talla</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="TallaBajo"
                  native-value="TallaBajo"
                  type="is-success"
                  >Menor o igual a 145cm</b-checkbox
                >
              </b-field>
            </td>
            <td>
              <b-field>
                <b-checkbox
                  v-model="TallaAlto"
                  native-value="TallaAlto"
                  type="is-danger"
                  >>145</b-checkbox
                >
              </b-field>
            </td>
            <td></td>
          </tr>

          <tr>
            <td>Esacolarida</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="EscolaridadBajo"
                  native-value="EscolaridadBajo"
                  type="is-success"
                  >Mayor a Primaria</b-checkbox
                >
              </b-field>
            </td>
            <td>
              <b-field>
                <b-checkbox
                  v-model="EscolaridadAlto"
                  native-value="EscolaridadAlto"
                  type="is-danger"
                  >Nula</b-checkbox
                >
              </b-field>
            </td>
            <td></td>
          </tr>
          <tr>
            <td>Trabajar más de 36hrs/10hrs al día de pie o sentada</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="TrabajarBajo"
                  native-value="TrabajarBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="TrabajarAlto"
                  native-value="TrabajarAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>
              Exposición a tóxicos(plomo, pesticidas, quimicos, Enf.Infecciosas)
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="ExposicionBajo"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="ExposicionAlto"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <th class="is-warning">
              <abbr title="Antecedentes"></abbr>ANTECEDENTES
            </th>
            <th class="is-warning">
              <abbr title="BajoRiesgo"></abbr>BAJO RIESGO
            </th>
            <th class="is-warning">
              <abbr title="AltoRiesgo" center></abbr>ALTO RIESGO
            </th>
            <th class="is-warning"><abbr title="AltoRiesgo" center></abbr></th>
          </tr>
          <tr>
            <td>Aborto previos Mayo o igual a 3(LUI)</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="AbortoBajo"
                  native-value="AbortoBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="AbortoAlto"
                  native-value="AbortoAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Partos prematuro (Menor a 28 SDG) previos menor o igual a 1</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="PartosBajo"
                  native-value="PartosBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="PartosAlto"
                  native-value="PartosAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Muertes perinatales u Óbitos</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="MuertesBajo"
                  native-value="MuertesBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="MuertesAlto"
                  native-value="MuertesAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Hijos con malformaciones</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="MalformacionBajo"
                  native-value="MarlformacionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="MalformacionAlto"
                  native-value="MarlformacionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Infección de vías urinarias(IVU Refactorio o recurrente)</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="IVUAlto"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >IVU baja</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="IVUBajo"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Piolonefritis</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Periodo intergenésico corto</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Mayor o igual a 2 años</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Menor a 12 meses</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Paridad</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >1 a 3</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Mayor o igual a4</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Cesárea anterior</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Mayor a 1</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>
              Parto anterior complicado (Hemorragia obstétrica, Retencion
              placentaria, TDPP)
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Sin complicaciones</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Complicado</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Preclampsia-eclampsia</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Productos previos</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Peso normal</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Menor a 2.5kg(restricción)/Mayor a
                  4hkg(macrosomico)</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>
              Cirugía genitourinaria(Miomectomía, conizacíon, Ectópico roto,
              etc.)
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Comorbilidades</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Ninguna</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <th class="is-warning">
              <abbr title="DatosGene"></abbr>EMBARAZO ACTUAL
            </th>
            <th class="is-warning">
              <abbr title="BajoRiesgo"></abbr>BAJO RIESGO
            </th>
            <th class="is-warning">
              <abbr title="AltoRiesgo" center></abbr>ALTO RIESGO
            </th>
            <th class="is-warning"><abbr title="AltoRiesgo" center></abbr></th>
          </tr>
          <tr>
            <td>Toxicomanias(Cocaína, tabaco, alcohol, crack, etc)</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Hepatitis B y C</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>
              "Rh (·} Profilaxis prenatal rutinaria con una administración de
              300 µg (1.500 UI) de inmunoalobulina anti-O, a la 24-30 SDG, 72
              hrs. Posterior Evento Obstétrico"
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>VIH(+)/SIFILIS(+)</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Hemoglobina en gr. 1er Trimestre del embarazo.</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >>11.5
                </b-checkbox>
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Menor a 9
                </b-checkbox>
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>
              Retardo menstrual +Dolor Abdominal (Descartar Embarazo Ectópico
              roto)
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Hiperémesis gravídica</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No presente</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Presente</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>
              Fondo Uterino/ SDG (Normal: 7-13 cm 1er;20- 26.8 cm 2do y 30-39 cm
              3er Trimestre) Descartar: Restricción; prematuro 6 Macrosomico;
              polidramnlos; Embarazo Molar 6 Múltiple
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Dentro Percentil (10 - 90)
                </b-checkbox>
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Por debajo de percentil 10 o por arriba del 90</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>
              "Descartar Diabetes gestacional a las 24-28 SDG (> 126 mg/dl
              ayuno, CTG 50gr. >200 ma/dl al Azar; > 200 ma/dl 2 horas (75 ar
              CTG)"
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>
              Mujeres con IMC >30 al inicio del embarazo, registre una ganancia
              de peso
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >4 A 9 Kg</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >> 9 Kg (IMC> 31.7 2DO Y > 33 3ER TRIMESTRE
                </b-checkbox>
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Amenaza de aborto/Menor a 500 gr o ≥ 22 SDG</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Infección Urinaria o Vaginosis</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Refactorio o recurrente</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>
              "Amenaza Parto Prematuro (22-36.6 SDG, contracciones, dilatación
              ≥3 cm._borramiento 90%, con o sin membranas amnióticas integras)"
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Enfermedad Hipertensiva del embarazo/ Preclamsia</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >TA 100/60 - 120/70</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >130/80 - 140 - 90
                </b-checkbox></b-field
              >
            </td>
            <td class="is-danger">
              <b-checkbox
                v-model="checkbox"
                native-value="ExposicionAlto"
                type="is-danger"
                >160-110 MMHG
              </b-checkbox>
            </td>
          </tr>
          <tr>
            <td>Proteinuria de Orina</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Normal</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >+ A + + +</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>
              Datos de severidad (cefalea, nauseas. dolor en epigastrio, edema,
              etc.)
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Ausentes</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Presentes</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Crisis Convulsivas</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Ausentes</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Presentes</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Malformaciones Conqénitas</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Determinar posición fetal</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Cefálico</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Pélvico, Transverso u oblicuo</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Movimientos fetales activos(+ 2 horas. después 28 SDG)</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Frecuencia Cardiaca Fetal</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >120 A 160 LAT X MIN</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Menor a 100 O >160 LATIDOS X MIN</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Implantación anómala de la Placenta</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si (Centarl al orificio cervical, marginal, acreta,
                  etc.)</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Hemorragia transvaginal (Obstétrica)</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si (Descartar DPPNI, Placenta previa)</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Taquicardia</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >100 a 119</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger">
              <b-checkbox
                v-model="checkbox"
                native-value="ExposicionAlto"
                type="is-danger"
                >Mayor o igual a 120 latidos X min</b-checkbox
              >
            </td>
          </tr>
          <tr>
            <td>Frecuencia Respiratoria</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >11 A 20 X MINUTO</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >21 A 30</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger">
              <b-checkbox
                v-model="checkbox"
                native-value="ExposicionAlto"
                type="is-danger"
                >≤ 10 Ó ≥ 30 X MIN</b-checkbox
              >
            </td>
          </tr>
          <tr>
            <td>Respuesta neurológica</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Alerta</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Voz</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Responde al dolor</b-checkbox
                >
              </b-field>
            </td>
          </tr>
          <tr>
            <td>Ruptura prematura de membranas</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Menor a 18 HORAS</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >≥ 18 HORAS</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Embarazo pos-termino 41 SDG sin trabajo de parto</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>Fiebre > 38 ºC</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>
              Síntomas Respiratorios (Tos seca, cefalea, anosmia, rinorrea,
              faringitis, etc.)
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >No</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Si</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
          <tr>
            <td>
              Concentración de Oxígeno /Dificultad para respirar, descartar
              Covid-19)
            </td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >90 a 100%</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Menos de 90</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-danger"></td>
          </tr>
          <tr>
            <td>Redes de apoyo</td>
            <td class="is-success">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionBajo"
                  type="is-success"
                  >Presente</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning">
              <b-field>
                <b-checkbox
                  v-model="checkbox"
                  native-value="ExposicionAlto"
                  type="is-danger"
                  >Ausente</b-checkbox
                >
              </b-field>
            </td>
            <td class="is-warning"></td>
          </tr>
        </tbody>
      </table>  
    </div> 
</template>
<script>
export default {
    data: ()=>(
        {
            checkbox: false,
            edadAlta:false,
            edadBajo:false,
            ImcBajo:false,
            ImcAlto:false,
            TallaBajo:false,
            TallaAlto:false,
            EscolaridadBajo:false,
            EscolaridadAlto:false,
            TrabajarBajo:false,
            TrabajarAlto:false,
            ExposicionBajo:false,
            ExposicionAlto:false,
            AbortoBajo:false,
            AbortoAlto:false,
            PartosBajo:false,
            PartosAlto:false,
            MuertesBajo:false,
            MuertesAlto:false,
            MuestesAlto:false,
            MalformacionBajo:false,
            MalformacionAlto:false,
            IVUAlto:false,
            IVUBajo:false,
          
         
    })
}
</script>