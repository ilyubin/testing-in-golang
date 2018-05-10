Feature: GET film

  Background:
    Given reset film

  Scenario Outline: Should return film
    When send get film request with <filmId>
    Then should receive film <title>

    Examples:
      | filmId | title                   |
      |     1  | A New Hope              |
      |     2  | The Empire Strikes Back |
      |     3  | Return of the Jedi      |
