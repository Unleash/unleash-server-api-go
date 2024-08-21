#!/bin/bash

operation_ids='["createApiToken", "updateApiToken", "deleteApiToken", "getApiTokensByName", "getAllApiTokens",
                "createUser", "updateUser", "getUser", "deleteUser",
                "getProjects", "createProject", "updateProject", "deleteProject",
                "getPermissions", "getProjectAccess", "setProjectAccess",
                "createRole", "getRoles", "getRoleById", "updateRole", "deleteRole",
                "getServiceAccounts", "createServiceAccount", "updateServiceAccount", "deleteServiceAccount",
                "getServiceAccountTokens", "createServiceAccountToken", "deleteServiceAccountToken"]'


jq --argjson operation_ids "$operation_ids" --argfile filters filter-ops.json '
  .paths |= with_entries(
    select(
      .value | any(
        .operationId as $id |
        $operation_ids | index($id)
      )
    )
  )

  |

 .components.schemas.projectSchema.properties |=
    with_entries(select(.key as $k | $filters.projectSchema | index($k) != null))

' modified-openapi.json > filtered-openapi.json