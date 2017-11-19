package lifegame;

import static org.junit.jupiter.api.Assertions.*;

import java.util.Arrays;
import java.util.List;

import org.junit.jupiter.api.Test;

import lifegame.impl2d.Field2d;

class GenerationTest {

    @Test
    void test() {
        final Field field = new Field2d(3, 3);
        final List<Status> statusList = Arrays.asList(
                Status.ALIVE, Status.ALIVE, Status.NOTHING,
                Status.ALIVE, Status.NOTHING, Status.NOTHING,
                Status.NOTHING, Status.NOTHING, Status.NOTHING);
        final Generation generation = new Generation(field, statusList);
        final Generation nextGeneration = generation.nextGeneration();
        final Generation expected = new Generation(field, Arrays.asList(
                Status.ALIVE, Status.ALIVE, Status.NOTHING,
                Status.ALIVE, Status.ALIVE, Status.NOTHING,
                Status.NOTHING, Status.NOTHING, Status.NOTHING));
        assertEquals(expected, nextGeneration);
    }
}
