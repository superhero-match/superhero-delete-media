/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package service

import (
	"github.com/superhero-match/superhero-delete-media/cmd/api/service/mapper"
)

// DeleteProfilePicture publishes message for a Superhero profile picture on Kafka topic for it to be
// consumed by consumer and deleted in DB, cache and Elasticsearch.
func (s *service) DeleteProfilePicture(superheroID string, position int64, deletedAt string) error {
	if err := s.Producer.DeleteProfilePicture(
		mapper.MapToProducerProfilePicture(
			superheroID,
			position,
			deletedAt,
		),
	); err != nil {
		return err
	}

	return nil
}
