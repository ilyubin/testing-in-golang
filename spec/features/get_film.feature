Feature: GET film

  Background:
    Given reset film

  Scenario Outline: Should return Luke Skywalker
    When send get film request with <filmId>
    Then should received film <title>

    Examples:
      | filmId | title                   |
      |     1  | A New Hope              |
      |     2  | The Empire Strikes Back |
      |     3  | Return of the Jedi      |
