package lifegame.impl2d;

import static java.util.function.Function.*;
import static java.util.stream.Collectors.*;
import static java.util.stream.IntStream.*;

import java.util.Collections;
import java.util.List;
import java.util.Objects;

import lifegame.Field;
import lifegame.Point;

public class Field2d implements Field {

    private final int width;
    private final int height;

    public Field2d(final int width, final int height) {
        this.width = width;
        this.height = height;
    }

    @Override
    public List<Point> getAroundPoints(final Point center) {
        if (center instanceof Point2d) {
            return ((Point2d) center).getAroundPoints(width, height);
        }
        return Collections.emptyList();
    }

    @Override
    public List<Point> getAllPoints() {
        return range(0, width)
                .mapToObj(x -> range(0, height).mapToObj(y -> new Point2d(x, y)))
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
        final Field2d other = (Field2d) obj;
        return width == other.width && height == other.height;
    }

    @Override
    public int hashCode() {
        return Objects.hash(width, height);
    }

    @Override
    public String toString() {
        return String.format("(%d, %d)", width, height);
    }
}
