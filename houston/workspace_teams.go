package houston

// AddTeamToWorkspace - add a team to a workspace
func (h ClientImplementation) AddWorkspaceTeam(workspaceUUID, teamUuid, role string) (*Workspace, error) {
	req := Request{
		Query:     WorkspaceTeamAddRequest,
		Variables: map[string]interface{}{"workspaceUuid": workspaceUUID, "teamUuid": teamUuid, "role": role},
	}

	r, err := req.DoWithClient(h.client)
	if err != nil {
		return nil, handleAPIErr(err)
	}

	return r.Data.AddWorkspaceTeam, nil
}

// RemoveTeamFromWorkspace - remove a team from a workspace
func (h ClientImplementation) DeleteWorkspaceTeam(workspaceUUID, teamUuid string) (*Workspace, error) {
	req := Request{
		Query:     WorkspaceTeamRemoveRequest,
		Variables: map[string]interface{}{"workspaceId": workspaceUUID, "teamUuid": teamUuid},
	}

	r, err := req.DoWithClient(h.client)
	if err != nil {
		return nil, handleAPIErr(err)
	}

	return r.Data.RemoveWorkspaceTeam, nil
}

// ListTeamAndRolesFromWorkspace - list teams and roles from a workspace
func (h ClientImplementation) ListWorkspaceTeamsAndRoles(workspaceUUID string) (*Workspace, error) {
	req := Request{
		Query:     WorkspacesGetRequest,
		Variables: map[string]interface{}{"workspaceId": workspaceUUID},
	}

	r, err := req.DoWithClient(h.client)
	if err != nil {
		return nil, handleAPIErr(err)
	}

	return &r.Data.GetWorkspaces[0], nil
}

// UpdateTeamRoleInWorkspace - update a team role in a workspace
func (h ClientImplementation) UpdateWorkspaceTeamRole(workspaceUUID, teamUuid, role string) (string, error) {
	req := Request{
		Query:     WorkspaceTeamUpdateRequest,
		Variables: map[string]interface{}{"workspaceUuid": workspaceUUID, "teamUuid": teamUuid, "role": role},
	}

	r, err := req.DoWithClient(h.client)
	if err != nil {
		return "", handleAPIErr(err)
	}

	return r.Data.WorkspaceUpdateTeamRole, nil
}

// GetTeamRoleInWorkspace - get a team role in a workspace
func (h ClientImplementation) GetWorkspaceTeamRole(workspaceUUID, teamUuid string) (WorkspaceTeamRoleBindings, error) {
	req := Request{
		Query:     WorkspaceGetTeamsRequest,
		Variables: map[string]interface{}{"workspaceUuid": workspaceUUID, "teamUuid": teamUuid},
	}

	r, err := req.DoWithClient(h.client)
	if err != nil {
		return WorkspaceTeamRoleBindings{}, handleAPIErr(err)
	}

	return r.Data.WorkspaceGetTeams, nil
}
