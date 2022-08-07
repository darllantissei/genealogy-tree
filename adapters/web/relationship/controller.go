package relationshipweb

import (
	"github.com/darllantissei/genealogy-tree/application/model"
)

func (pw *RelationshipWeb) parseDataToApp(relationshipDTO RelationshipDTO) model.Relationship {

	rtshpApp := model.Relationship{
		ID:       relationshipDTO.ID,
		PersonID: relationshipDTO.PersonID,
	}

	for _, memberDTO := range relationshipDTO.Members {
		rtshpApp.Members = append(rtshpApp.Members, model.RelationshipMember{
			PersonID:       memberDTO.PersonID,
			RelationshipID: memberDTO.RelationshipID,
			Type:           memberDTO.Type,
			Kindship:       memberDTO.Kindship,
		})
	}

	return rtshpApp
}

func (pw *RelationshipWeb) parseData(rtshpApp model.Relationship) RelationshipDTO {

	rtshpDTO := RelationshipDTO{
		ID:       rtshpApp.ID,
		PersonID: rtshpApp.PersonID,
	}

	for _, memberApp := range rtshpApp.Members {

		rtshpDTO.Members = append(rtshpDTO.Members, RelationshipMember{
			PersonID:       memberApp.PersonID,
			RelationshipID: memberApp.RelationshipID,
			Type:           memberApp.Type,
			Kindship:       memberApp.Kindship,
		})
	}

	return rtshpDTO
}
