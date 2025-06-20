# Contribuciones de Ariana López

## Sprint 1 - Estructura inicial iac

- 2025-06-9: Configuré la estructura inicial del directorio iac/ para gestionar infraestructura como código con Terraform  
  Commit: 
  - feat(iac): definir variables iniciales
  - feat(iac): definir recursos dummy
  - feat(iac): establecer etiqueta obligatoria

  PR: #7

- 2025-06-11: Creé el script run_all_scans.sh que orquesta la ejecución de todos los escáneres de seguridad.  
  Commit: 
  - feat(scan): agregar script que ejecuta todos los scans
  - doc(readme): actualizar doc de script bash

  PR: #5

## Sprint 2 - Configuración de herramientas de seguridad

- 2025-06-15: Refiné la salida de Bandit para excluir reglas y mostrar solo vulnerabilidades críticas.  
  Commit: 
  - chore(bandit): filtrar solo vulnerabilidades críticas
  - chore(bandit): excluir tests específicos del escaneo

  PR: #23

- 2025-06-17: Configuré Checkov con un ruleset para validar etiquetas obligatorias en recursos Terraform.  
  Commit: 
  - feat(checkov): agregar ruleset personalizado para validación de etiquetas
  - chore(checkov): habilitar ejecución y filtrado de resultados del ruleset personalizado
  - feat(iac): agregar recursos dummy con skip de ruleset personalizado

  PR: #23

## Sprint 3 - Validación de configuración de red

- 2025-06-19: Implementé el módulo network_dummy que genera un archivo network_config.json con parámetros de red.  
  Commit: 
  - feat(iac): añadir módulo network_dummy para generar archivo JSON de configuración de red  
  
  PR: #27

- 2025-06-19: Integré validación de red en scan_checkov.sh para analizar network_config.json.  
  Commit: 
  - feat(security): integrar validación de red al reporte de seguridad general
  - feat(scan): integrar validación de configuración de red
  - fix(scan): corregir ruta de salida para resultados por módulo en scan_tflint

  PR: #27