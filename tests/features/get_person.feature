Feature: GET person

  Scenario: Should return Luke Skywalker
    When send get person request with personId 1
    Then should received person "Luke Skywalker"
    And it hair is "blond"
    And it eye is "blue"

  Scenario: Should return R2-D2
    When send get person request with personId 3
    Then should received person "R2-D2"
    And it eye is "red"

  Scenario: Should be not found error
    When send get person request with personId 0 expected error 404
    Then should be error "Not found"
