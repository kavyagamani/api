Feature: Map redirect API

  Background:
    * url baseUrl

  Scenario: Redirect to Google Maps when location provided
    Given path 'map'
    And param location = 'New York'
    When method get
    Then status 302
    And header Location contains 'https://www.google.com/maps'

  Scenario: Missing location returns 400 JSON error
    Given path 'map'
    When method get
    Then status 400
    And match response == { code: '#string', message: '#string' }

  Scenario: Root serves frontend index (health/basic availability)
    Given path ''
    When method get
    Then status 200
    And match response contains '<html'
