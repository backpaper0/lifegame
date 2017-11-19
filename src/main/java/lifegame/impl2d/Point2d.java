package lifegame.impl2d;

import static java.util.function.Function.*;
import static java.util.stream.Collectors.*;
import static java.util.stream.IntStream.*;

import java.util.List;
import java.util.Objects;

import lifegame.Point;

public final class Point2d implements Point {

    private final int x;
    private final int y;

    public Point2d(final int x, final int y) {
        this.x = x;
        this.y = y;
    }

    List<Point> getAroundPoints(final int width, final int height) {
        return rangeClosed(x - 1, x + 1)
                .filter(a -> 0 <= a && a < width)
                .mapToObj(
                        a -> rangeClosed(y - 1, y + 1).filter(b -> 0 <= b && b < height)
                                .filter(b -> (a == x && b == y) == false)
                                .mapToObj(b -> new Point2d(a, b)))
                .flatMap(identity())
                .collect(toList());
    }

    @Override
    public boolean equals(final Object obj) {
        if (this == obj) {
            return true;
        } else if (obj == null) {
            return false;
        } else if (obj.getClass() != getClass()) {
            return false;
        }
        final Point2d other = (Point2d) obj;
        return x == other.x && y == other.y;
    }

    @Override
    public int hashCode() {
        return Objects.hash(x, y);
    }

    @Override
    public String toString() {
        return String.format("(%d, %d)", x, y);
    }
}
