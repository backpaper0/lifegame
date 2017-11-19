package lifegame;

import static org.junit.jupiter.api.Assertions.*;

import java.util.stream.Stream;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.MethodSource;

class StatusTest {

    @ParameterizedTest
    @MethodSource("fixturesForNoting")
    void NOTHING(final Fixture fixture) {
        final Status nextStatus = Status.NOTHING.nextStatus(fixture.aliveCount);
        assertEquals(fixture.expected, nextStatus);
    }

    @ParameterizedTest
    @MethodSource("fixturesForAlive")
    void alive(final Fixture fixture) {
        final Status nextStatus = Status.ALIVE.nextStatus(fixture.aliveCount);
        assertEquals(fixture.expected, nextStatus);
    }

    static Stream<Fixture> fixturesForNoting() {
        return Stream.of(
                new Fixture(0, Status.NOTHING),
                new Fixture(1, Status.NOTHING),
                new Fixture(2, Status.NOTHING),
                new Fixture(3, Status.ALIVE),
                new Fixture(4, Status.NOTHING),
                new Fixture(5, Status.NOTHING),
                new Fixture(6, Status.NOTHING),
                new Fixture(7, Status.NOTHING),
                new Fixture(8, Status.NOTHING));
    }

    static Stream<Fixture> fixturesForAlive() {
        return Stream.of(
                new Fixture(0, Status.NOTHING),
                new Fixture(1, Status.NOTHING),
                new Fixture(2, Status.ALIVE),
                new Fixture(3, Status.ALIVE),
                new Fixture(4, Status.NOTHING),
                new Fixture(5, Status.NOTHING),
                new Fixture(6, Status.NOTHING),
                new Fixture(7, Status.NOTHING),
                new Fixture(8, Status.NOTHING));
    }

    static class Fixture {
        private final int aliveCount;
        private final Status expected;

        public Fixture(final int aliveCount, final Status expected) {
            this.aliveCount = aliveCount;
            this.expected = expected;
        }

        @Override
        public String toString() {
            return String.format("%d %s", aliveCount, expected);
        }
    }
}
