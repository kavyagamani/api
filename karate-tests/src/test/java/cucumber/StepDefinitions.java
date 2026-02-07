package cucumber;

import static org.junit.jupiter.api.Assertions.assertEquals;

import io.cucumber.java.en.Given;
import io.cucumber.java.en.Then;
import io.cucumber.java.en.When;

public class StepDefinitions {

    private int a;
    private int b;
    private int result;

    @Given("a number {int}")
    public void a_number(int num) {
        this.a = num;
    }

    @When("I add {int}")
    public void i_add(int num) {
        this.b = num;
        this.result = this.a + this.b;
    }

    @Then("the result should be {int}")
    public void the_result_should_be(int expected) {
        assertEquals(expected, this.result);
    }
}
